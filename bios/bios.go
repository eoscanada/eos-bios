package bios

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash/crc64"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/eoscanada/eos-bios/bios/disco"
	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
)

type BIOS struct {
	Network *Network
	// MyPeers represent the peers my local node will handle. It is
	// plural because when launching a 3-node network, your peer will
	// be cloned a few time to have a full schedule of 21 producers.
	MyPeers []*Peer

	SingleOnly               bool
	OverrideBootSequenceFile string
	Log                      *Logger

	LaunchDisco  *disco.Discovery
	TargetNetAPI *eos.API
	Snapshot     Snapshot
	BootSequence []*OperationType

	Genesis *GenesisJSON

	// ShuffledProducers is an ordered list of producers according to
	// the shuffled peers.
	RandSource        rand.Source
	ShuffledProducers []*Peer

	EphemeralPrivateKey *ecc.PrivateKey
	EphemeralPublicKey  ecc.PublicKey
}

func NewBIOS(logger *Logger, network *Network, targetAPI *eos.API) *BIOS {
	network.Log = logger
	b := &BIOS{
		Network:      network,
		TargetNetAPI: targetAPI,
		Log:          logger,
	}
	return b
}

func (b *BIOS) SetGenesis(gen *GenesisJSON) {
	b.Genesis = gen
}

func (b *BIOS) Init() error {
	// HAVE TWO init sequences:
	// * a first that does BASIC inits
	//
	// * a second that can be recalled just after receiving the Launch
	// Block (and everyone fell in agreement), which might change the
	// bootsequence we've agreed upon, might change the network,
	// topology, etc..  download all the contents.. because topology

	// Load launch data
	launchDisco, err := b.Network.ConsensusDiscovery()
	if err != nil {
		return fmt.Errorf("couldn'get consensus on launch data: %s", err)
	}

	b.LaunchDisco = launchDisco

	// FIXME: we should call `setProducers()` after a call to `waitLaunchBlock()`, or call it again
	// now that we have the shuffling ready..
	if err := b.setProducers(); err != nil {
		return err
	}

	if err = b.setMyPeers(); err != nil {
		return fmt.Errorf("error setting my producer definitions: %s", err)
	}

	// Load Boot Sequence...
	bootseqFile, err := b.GetContentsCacheRef("boot_sequence.yaml")
	if err != nil {
		return err
	}

	var rawBootSeq []byte
	if b.OverrideBootSequenceFile != "" {
		b.Log.Printf("Using overridden boot sequence from %q\n", b.OverrideBootSequenceFile)

		rawBootSeq, err = ioutil.ReadFile(b.OverrideBootSequenceFile)
		if err != nil {
			return fmt.Errorf("reading overridden boot_sequence file: %s", err)
		}
	} else {
		rawBootSeq, err = b.Network.ReadFromCache(bootseqFile)
		if err != nil {
			return fmt.Errorf("reading boot_sequence file: %s", err)
		}
	}

	var bootSeq struct {
		BootSequence []*OperationType `json:"boot_sequence"`
	}
	if err := yamlUnmarshal(rawBootSeq, &bootSeq); err != nil {
		return fmt.Errorf("loading boot sequence: %s", err)
	}

	// TODO: we need to RELOAD the boot sequence from the selected
	// decided upon, once the Launch Block is reached.
	b.BootSequence = bootSeq.BootSequence

	return nil
}

func (b *BIOS) StartOrchestrate() error {
	b.Log.Println("Starting Orchestraion process", time.Now())

	b.Log.Println("Showing pre-randomized network discovered:")
	b.PrintProducerSchedule()

	b.RandSource = b.waitLaunchBlock()

	// Randomize the list now.
	if err := b.setProducers(); err != nil {
		return err
	}

	// Once we have it, we can discover the net again (unless it's been discovered VERY recently)
	// and we b.Init() again.. so load the latest version of the LaunchData according to this
	// potentially new discovery network.
	b.Log.Println("Seed network block used to seed randomization, updating graph one last time...")

	if err := b.Network.UpdateGraph(); err != nil {
		return fmt.Errorf("update graph: %s", err)
	}

	b.Log.Println("Network used for launch:")
	b.PrintProducerSchedule()

	if err := b.DispatchInit("orchestrate"); err != nil {
		return fmt.Errorf("dispatch init hook: %s", err)
	}

	switch b.MyRole() {
	case RoleBootNode:
		if err := b.RunBootSequence(); err != nil {
			return fmt.Errorf("as boot node: %s", err)
		}
	case RoleABP:
		if err := b.RunJoinNetwork(true, true); err != nil {
			return fmt.Errorf("as abp: %s", err)
		}
	default:
		if err := b.RunJoinNetwork(true, false); err != nil {
			return fmt.Errorf("as participant: %s", err)
		}
	}

	return b.DispatchDone("orchestrate")
}

func (b *BIOS) StartJoin(validate bool) error {
	b.Log.Println("Starting network join process", time.Now())

	b.PrintProducerSchedule()

	if err := b.DispatchInit("join"); err != nil {
		return fmt.Errorf("dispatch init hook: %s", err)
	}

	if err := b.RunJoinNetwork(validate, false); err != nil {
		return fmt.Errorf("join network: %s", err)
	}

	return b.DispatchDone("join")
}

func (b *BIOS) StartBoot() error {
	b.Log.Println("Starting network join process", time.Now())

	b.PrintProducerSchedule()

	if err := b.DispatchInit("boot"); err != nil {
		return fmt.Errorf("dispatch init hook: %s", err)
	}

	if err := b.RunBootSequence(); err != nil {
		return fmt.Errorf("run bios boot: %s", err)
	}

	return b.DispatchDone("boot")
}

func (b *BIOS) PrintProducerSchedule() {
	b.Network.PrintOrderedPeers()

	b.Log.Println("")
	b.Log.Println("###############################################################################################")
	b.Log.Println("")
	if b.AmIBootNode() {
		b.Log.Println("                              MY ROLE: BIOS BOOT NODE")
	} else if b.AmIAppointedBlockProducer() {
		b.Log.Println("                              MY ROLE: APPOINTED BLOCK PRODUCER")
	} else {
		b.Log.Println("                              MY ROLE: JOINING NETWORK")
	}
	b.Log.Println("")

	b.Log.Println("###############################################################################################")
	b.Log.Println("")
}

func (b *BIOS) RunBootSequence() error {
	b.Log.Println("START BOOT SEQUENCE...")

	// keys, _ := b.TargetNetAPI.Signer.(*eos.KeyBag).AvailableKeys()
	// for _, key := range keys {
	// 	b.Log.Println("Available key in the KeyBag:", key)
	// }

	// Update boot sequence HERE

	ephemeralPrivateKey, err := b.GenerateEphemeralPrivKey()
	if err != nil {
		return err
	}

	b.EphemeralPrivateKey = ephemeralPrivateKey
	pubKey := ephemeralPrivateKey.PublicKey()
	b.EphemeralPublicKey = pubKey

	// b.TargetNetAPI.Debug = true

	privKey := ephemeralPrivateKey.String()

	b.Log.Printf("Generated ephemeral keys:\n\n\tPublic key: %s\n\tPrivate key: %s..%s\n\n", pubKey, privKey[:7], privKey[len(privKey)-7:])

	// Store keys in wallet, to sign `SetCode` and friends..
	if err := b.TargetNetAPI.Signer.ImportPrivateKey(privKey); err != nil {
		return fmt.Errorf("ImportWIF: %s", err)
	}

	genesisData := b.GenerateGenesisJSON(pubKey.String())

	if len(b.Network.MyPeer.Discovery.SeedNetworkPeers) > 0 && !b.SingleOnly {

		b.Log.Printf("Publishing genesis data to the seed network... ")
		_, err := b.Network.SeedNetAPI.SignPushActions(
			disco.NewUpdateGenesis(b.Network.MyPeer.Discovery.SeedNetworkAccountName, genesisData, []string{}),
		)
		if err != nil {
			b.Log.Println("")
			return fmt.Errorf("updating genesis on seednet: %s", err)
		}
		b.Log.Println(" done")

		if err = b.DispatchBootPublishGenesis(genesisData); err != nil {
			return fmt.Errorf("dispatch boot_publish_genesis hook: %s", err)
		}
	}

	orderedPeers := b.Network.OrderedPeers(b.Network.MyNetwork())
	otherPeers := b.someTopmostPeersAddresses(orderedPeers)
	if err := b.DispatchBootNode(genesisData, pubKey.String(), privKey, otherPeers); err != nil {
		return fmt.Errorf("dispatch boot_node hook: %s", err)
	}

	b.pingTargetNetwork()

	b.Log.Println("In-memory keys:")
	memkeys, _ := b.TargetNetAPI.Signer.AvailableKeys()
	for _, key := range memkeys {
		b.Log.Printf("- %s\n", key.String())
	}
	b.Log.Println("")

	//eos.Debug = true

	for _, step := range b.BootSequence {
		b.Log.Printf("%s  [%s] ", step.Label, step.Op)

		if b.LaunchDisco.TargetNetworkIsTest == 0 {
			step.Data.ResetTestnetOptions()
		}

		acts, err := step.Data.Actions(b)
		if err != nil {
			return fmt.Errorf("getting actions for step %q: %s", step.Op, err)
		}

		if len(acts) != 0 {
			for idx, chunk := range ChunkifyActions(acts) {
				//time.Sleep(500 * time.Millisecond)
				err := Retry(25, time.Second, func() error {
					_, err := b.TargetNetAPI.SignPushActions(chunk...)
					if err != nil {
						b.Log.Printf("r")
						b.Log.Debugf("error pushing transaction for step %q, chunk %d: %s\n", step.Op, idx, err)
						return fmt.Errorf("push actions for step %q, chunk %d: %s", step.Op, idx, err)
					}
					return nil
				})
				if err != nil {
					b.Log.Printf(" failed\n")
					return err
				}
				b.Log.Printf(".")
			}
			b.Log.Printf(" done\n")
		}
	}

	b.Log.Println("Waiting 2 seconds for transactions to flush to blocks")
	time.Sleep(2 * time.Second)

	// FIXME: don't do chain validation here..
	isValid, err := b.RunChainValidation()
	if err != nil {
		return fmt.Errorf("chain validation: %s", err)
	}
	if !isValid {
		b.Log.Println("WARNING: chain invalid, destroying network if possible")
		os.Exit(0)
	}

	if err := b.DispatchBootPublishHandoff(); err != nil {
		return fmt.Errorf("dispatch boot_publish_handoff: %s", err)
	}

	return nil
}

func (b *BIOS) getMyPeerVariations() (out []*Peer) {
	for _, prod := range b.ShuffledProducers {
		if prod.Discovery.SeedNetworkAccountName == b.Network.MyPeer.Discovery.SeedNetworkAccountName {
			out = append(out, prod)
		}
	}
	return
}

func (b *BIOS) RunJoinNetwork(validate, sabotage bool) error {
	if b.Genesis == nil {
		if b.SingleOnly {
			b.Genesis = b.inputGenesisData()
		} else {
			b.Genesis = b.pollGenesisData()
		}
	}

	pubKey, err := ecc.NewPublicKey(b.Genesis.InitialKey)
	if err != nil {
		return fmt.Errorf("invalid genesis public key: %s", err)
	}
	b.EphemeralPublicKey = pubKey

	// Create mesh network
	otherPeers := b.computeMyMeshP2PAddresses()

	if err := b.DispatchJoinNetwork(b.Genesis, b.getMyPeerVariations(), otherPeers); err != nil {
		return fmt.Errorf("dispatch join_network hook: %s", err)
	}

	if validate {
		b.Log.Println("###############################################################################################")
		b.Log.Println("Launching chain validation")

		isValid, err := b.RunChainValidation()
		if err != nil {
			return fmt.Errorf("chain validation: %s", err)
		}
		if !isValid {
			b.Log.Println("WARNING: chain invalid, destroying network if possible")
			os.Exit(0)
		}
	} else {
		b.Log.Println("")
		b.Log.Println("Not doing chain validation. Someone else will do it.")
		b.Log.Println("")
	}

	// TODO: loop operations, check all actions against blocks that you can fetch from here.
	// Check ALL actions, should match the orchestrated launch data:
	// - otherwise, sabotage

	if validate {
		b.waitOnHandoff(b.Genesis)
	}

	return nil
}

func (b *BIOS) RunChainValidation() (bool, error) {
	bootSeqMap := ActionMap{}
	bootSeq := []*eos.Action{}

	for _, step := range b.BootSequence {
		if b.LaunchDisco.TargetNetworkIsTest == 0 {
			step.Data.ResetTestnetOptions()
		}

		acts, err := step.Data.Actions(b)
		if err != nil {
			return false, fmt.Errorf("validating: getting actions for step %q: %s", step.Op, err)
		}

		for _, stepAction := range acts {
			if stepAction == nil {
				continue
			}

			data, err := eos.MarshalBinary(stepAction)
			if err != nil {
				return false, fmt.Errorf("validating: binary marshalling: %s", err)
			}
			stepAction.SetToServer(false)
			key := sha2(data)

			// if _, ok := bootSeqMap[key]; ok {
			// 	// TODO: don't fatal here plz :)
			// 	log.Fatalf("Same action detected twice [%s] with key [%s]\n", stepAction.Name, key)
			// }
			bootSeqMap[key] = stepAction
			bootSeq = append(bootSeq, stepAction)
		}

	}

	err := b.validateTargetNetwork(bootSeqMap, bootSeq)
	if err != nil {
		b.Log.Printf("BOOT SEQUENCE VALIDATION FAILED:\n%s", err)
		return false, nil
	}

	b.Log.Println("")
	b.Log.Println("All good! Chain validation succeeded!")
	b.Log.Println("")

	return true, nil
}

type ActionMap map[string]*eos.Action

type ValidationError struct {
	Err               error
	BlockNumber       int
	Action            *eos.Action
	RawAction         []byte
	Index             int
	ActionHexData     string
	PackedTransaction eos.PackedTransaction
}

func (e ValidationError) Error() string {
	s := fmt.Sprintf("Action [%d][%s::%s] from block not found in boot sequences\n", e.Index, e.Action.Account, e.Action.Name)

	data, err := json.Marshal(e.Action)
	if err != nil {
		s += fmt.Sprintf("    json generation err: %s\n", err)
	} else {
		s += fmt.Sprintf("    json data: %s\n", string(data))
	}
	s += fmt.Sprintf("    hex data: %s\n", hex.EncodeToString(e.RawAction))
	s += fmt.Sprintf("    error: %s\n", e.Err.Error())

	return s
}

type ValidationErrors struct {
	Errors []error
}

func (v ValidationErrors) Error() string {

	s := ""
	for _, err := range v.Errors {
		s += ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n"
		s += err.Error()
		s += "<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<\n"
	}

	return s
}

func (b *BIOS) pingTargetNetwork() {
	b.Log.Printf("Pinging target network at %q...", b.TargetNetAPI.BaseURL)
	for {
		info, err := b.TargetNetAPI.GetInfo()
		if err != nil {
			b.Log.Debugf("target network error: %s\n", err)
			b.Log.Printf("e")
			time.Sleep(1 * time.Second)
			continue
		}

		if info.HeadBlockNum < 2 {
			b.Log.Debugln("target network: still no blocks in")
			b.Log.Printf(".")
			time.Sleep(1 * time.Second)
			continue
		}

		break
	}

	b.Log.Println(" touchdown!")

}

func (b *BIOS) validateTargetNetwork(bootSeqMap ActionMap, bootSeq []*eos.Action) (err error) {
	expectedActionCount := len(bootSeq)
	validationErrors := make([]error, 0)

	b.pingTargetNetwork()

	// TODO: wait for target network to be up, and responding...
	b.Log.Println("Pulling blocks from chain until we gathered all actions to validate:")
	lastDisplay := time.Now()
	blockHeight := 1
	actionsRead := 0
	seenMap := map[string]bool{}

	for {
		m, err := b.TargetNetAPI.GetBlockByNum(uint32(blockHeight))
		if err != nil {
			b.Log.Debugln("Failed getting block num from target api:", err)
			b.Log.Printf("e")
			time.Sleep(1 * time.Second)
			continue
		} else {
			b.Log.Printf(".\n")
		}

		// TODO: Once we hit a failure, we should now poke them each 500ms, and try to speed up slllightly

		if time.Now().Sub(lastDisplay) > 1*time.Second {
			b.Log.Printf("Block %d, read actions %d/%d\n", blockHeight, actionsRead, len(bootSeq))
			lastDisplay = time.Now()
		}

		blockHeight++

		b.Log.Printf("Receiving block tmroot=%q producer=%s transactions=%d\n", hex.EncodeToString(m.TransactionMRoot), m.Producer, len(m.Transactions))

		for _, receipt := range m.Transactions {
			unpacked, err := receipt.Transaction.Packed.Unpack()
			if err != nil {
				b.Log.Println("WARNING: Unable to unpack transaction, won't be able to fully validate:", err)
				return fmt.Errorf("unpack transaction failed")
			}

			for _, act := range unpacked.Actions {
				act.SetToServer(false)
				data, err := eos.MarshalBinary(act)
				if err != nil {
					b.Log.Printf("Error marshalling an action: %s\n", err)
					validationErrors = append(validationErrors, ValidationError{
						Err:               err,
						BlockNumber:       1, // extract from the block transactionmroot
						PackedTransaction: receipt.Transaction.Packed,
						Action:            act,
						RawAction:         data,
						ActionHexData:     hex.EncodeToString(act.HexData),
						Index:             actionsRead,
					})
					return err
				}
				key := sha2(data) // TODO: compute a hash here..

				b.Log.Printf("- Validating action %d/%d [%s::%s]", actionsRead+1, expectedActionCount, act.Account, act.Name)
				if _, ok := bootSeqMap[key]; !ok {
					validationErrors = append(validationErrors, ValidationError{
						Err:               errors.New("not found"),
						BlockNumber:       1, // extract from the block transactionmroot
						PackedTransaction: receipt.Transaction.Packed,
						Action:            act,
						RawAction:         data,
						ActionHexData:     hex.EncodeToString(act.HexData),
						Index:             actionsRead,
					})
					b.Log.Printf(" INVALID ***************************** INVALID *************.\n")
				} else {
					seenMap[key] = true
					b.Log.Printf(" valid.\n")
				}

				actionsRead++
			}
		}

		if actionsRead == len(bootSeq) {
			break
		}

	}

	if len(validationErrors) > 0 {
		b.Log.Println("Unseen transactions:")
		for _, act := range bootSeq {
			data, _ := eos.MarshalBinary(act)
			key := sha2(data)
			if !seenMap[key] {
				b.Log.Printf("- Action %s::%s [%q]\n", act.Account, act.Name, hex.EncodeToString(data))
			}
		}
		return ValidationErrors{Errors: validationErrors}
	}

	return nil
}

func (b *BIOS) waitLaunchBlock() rand.Source {
	targetBlockNum := uint32(b.LaunchDisco.SeedNetworkLaunchBlock)

	b.Log.Println("Polling seed network until launch block, target:", targetBlockNum)

	for {
		launchTime, _, err := b.Network.LaunchBlockTime(targetBlockNum)
		if err != nil {
			b.Log.Println(err.Error())
		}

		if launchTime.After(time.Now()) {
			b.Log.Printf("- not yet, %s to go\n", launchTime.Sub(time.Now()))
			time.Sleep(time.Second)
			continue
		}

		hash, err := b.Network.GetBlockHeight(targetBlockNum)
		if err != nil {
			b.Log.Println("error fetching seed network's target block hash:", err)
			time.Sleep(2 * time.Second)
			continue
		}

		b.Log.Println("- got block", targetBlockNum, "- hash is", hex.EncodeToString(hash))
		chksum := crc64.Checksum(hash, crc64.MakeTable(crc64.ECMA))
		return rand.NewSource(int64(chksum))

	}
}

func (b *BIOS) pollGenesisData() (genesis *GenesisJSON) {
	b.Log.Println("")
	b.Log.Println("Waiting for the BIOS Boot node to publish the genesis data to the seed network contract..")

	bootNode := b.ShuffledProducers[0]

	b.Log.Printf("Polling account %q...", bootNode.Discovery.SeedNetworkAccountName)
	for {
		time.Sleep(500 * time.Millisecond)

		b.Log.Printf(".")
		genesisData, err := b.Network.PollGenesisTable(bootNode.Discovery.SeedNetworkAccountName)
		if err != nil {
			b.Log.Debugf("\n- data not ready: %s", err)
			continue
		}

		if len(genesisData) == 0 {
			b.Log.Debugf("\n- data still empty")
			continue
		}

		err = json.Unmarshal([]byte(genesisData), &genesis)
		if err != nil {
			b.Log.Debugf("\n- data not valid: %q (err=%s)", err, genesisData)
			continue
		}

		b.Log.Println("")
		b.Log.Println("Got genesis data:")
		b.Log.Println("    ", genesisData)
		b.Log.Println("")
		b.Log.Printf("    Public key for new launch: %s\n", genesis.InitialKey)
		b.Log.Println("")

		return
	}
}

func (b *BIOS) inputGenesisData() (genesis *GenesisJSON) {
	b.Log.Println("")

	for {
		b.Log.Printf("Please input the genesis data of the network you want to join: ")
		genesisData, err := ScanSingleLine()
		if err != nil {
			b.Log.Println("error reading:", err)
			continue
		}

		err = json.Unmarshal([]byte(genesisData), &genesis)
		if err != nil {
			b.Log.Printf("Invalid genesis data: %s\n", err)
			continue
		}

		return
	}
}

func (b *BIOS) waitOnHandoff(genesis *GenesisJSON) {
	b.Log.Println("------------------")
	b.Log.Println("This step is to prove the BIOS Boot node kept nothing to itself.")

	for {
		b.Log.Printf("Please paste the EPHEMERAL private key that the Boot node published: ")
		privKey, err := ScanSingleLine()
		if err != nil {
			b.Log.Println("Error reading line:", err)
			continue
		}

		privKey = strings.TrimSpace(privKey)

		key, err := ecc.NewPrivateKey(privKey)
		if err != nil {
			b.Log.Println("Invalid private key pasted:", err)
			continue
		}

		if key.PublicKey().String() == genesis.InitialKey {
			b.Log.Println("")
			b.Log.Println("   HANDOFF VERIFIED! EOS CHAIN IS ALIVE !")
			b.Log.Println("")
			return
		} else {
			b.Log.Println("")
			b.Log.Println("   WARNING: private key provided does NOT match the genesis data")
			b.Log.Println("")
		}
	}
}

func (b *BIOS) GenerateEphemeralPrivKey() (*ecc.PrivateKey, error) {
	return ecc.NewRandomPrivateKey()
}

func (b *BIOS) GenerateGenesisJSON(pubKey string) string {
	// known not to fail
	cnt, _ := json.Marshal(&GenesisJSON{
		InitialTimestamp: time.Now().UTC().Format("2006-01-02T15:04:05"),
		InitialKey:       pubKey,
	})
	return string(cnt)
}

func (b *BIOS) GetContentsCacheRef(filename string) (string, error) {
	for _, fl := range b.LaunchDisco.TargetContents {
		if fl.Name == filename {
			return fl.Ref, nil
		}
	}
	return "", fmt.Errorf("%q not found in target contents", filename)
}

func (b *BIOS) setProducers() error {
	network := b.Network.MyNetwork()
	orderedPeers := b.Network.OrderedPeers(network)

	b.ShuffledProducers = orderedPeers

	b.shuffleProducers() // conditionally

	// We'll multiply the other producers as to have a full schedule
	if len(b.ShuffledProducers) > 1 {
		if numProds := len(b.ShuffledProducers); numProds < 22 {
			cloneCount := numProds - 1
			count := 0
			for {
				if len(b.ShuffledProducers) == 22 {
					break
				}

				fromPeer := b.ShuffledProducers[1+count%cloneCount]
				count++

				clonedDisco := *fromPeer.Discovery

				accountVar := eos.AccountName("")
				for {
					accountVar = accountVariation(fromPeer.Discovery.TargetAccountName, count)
					if b.targetInShuffledProducers(accountVar) {
						count++
						continue
					}
					break
				}
				clonedDisco.TargetAccountName = accountVar
				clonedPeer := &Peer{
					Discovery: &clonedDisco,
					UpdatedAt: fromPeer.UpdatedAt,
				}
				b.ShuffledProducers = append(b.ShuffledProducers, clonedPeer)
			}
		}
	}

	return nil
}

func (b *BIOS) targetInShuffledProducers(acct eos.AccountName) bool {
	for _, prod := range b.ShuffledProducers {
		if prod.Discovery.TargetAccountName == acct {
			return true
		}
	}
	return false
}

func (b *BIOS) shuffleProducers() {
	if b.RandSource == nil {
		b.Log.Println("Random source not set, skipping producer shuffling")
		return
	}

	b.Log.Println("Shuffling producers listed in the launch file")
	r := rand.New(b.RandSource)
	// shuffle top 25%, capped to 5
	shuffleHowMany := int64(math.Min(math.Ceil(float64(len(b.ShuffledProducers))*0.25), 5))
	if shuffleHowMany > 1 {
		b.Log.Println("- Shuffling top", shuffleHowMany)
		for round := 0; round < 100; round++ {
			from := r.Int63() % shuffleHowMany
			to := r.Int63() % shuffleHowMany
			if from == to {
				continue
			}

			//b.Log.Println("Swapping from", from, "to", to)
			b.ShuffledProducers[from], b.ShuffledProducers[to] = b.ShuffledProducers[to], b.ShuffledProducers[from]
		}
	} else {
		b.Log.Println("- No shuffling, network too small")
	}
}

func (b *BIOS) IsBootNode(account string) bool {
	return string(b.ShuffledProducers[0].AccountName()) == account
}

func (b *BIOS) AmIBootNode() bool {
	return b.IsBootNode(string(b.Network.MyPeer.Discovery.TargetAccountName))
}

func (b *BIOS) MyRole() Role {
	if b.AmIBootNode() {
		return RoleBootNode
	} else if b.AmIAppointedBlockProducer() {
		return RoleABP
	}
	return RoleParticipant
}

func (b *BIOS) IsAppointedBlockProducer(account string) bool {
	for i := 1; i < 22 && len(b.ShuffledProducers) > i; i++ {
		if string(b.ShuffledProducers[i].Discovery.TargetAccountName) == account {
			return true
		}
	}
	return false
}

func (b *BIOS) AmIAppointedBlockProducer() bool {
	return b.IsAppointedBlockProducer(string(b.Network.MyPeer.Discovery.TargetAccountName))
}

// MyProducerDefs will provide more than one producer def ONLY when
// your launch files contains LESS than 21 potential appointed block
// producers.  This way, you can have your nodes respond to many
// account names and have the network function. Your producer will
// simply produce more blocks, under different names.
func (b *BIOS) setMyPeers() error {
	myPeer := b.Network.MyPeer

	out := []*Peer{myPeer}

	for _, peer := range b.ShuffledProducers {
		if peer.Discovery.TargetAccountName == myPeer.Discovery.TargetAccountName {
			out = append(out, peer)
		}
	}

	b.MyPeers = out

	return nil
}

func ChunkifyActions(actions []*eos.Action) (out [][]*eos.Action) {
	currentChunk := []*eos.Action{}
	for _, act := range actions {
		if act == nil {
			if len(currentChunk) != 0 {
				out = append(out, currentChunk)
			}
			currentChunk = []*eos.Action{}
		} else {
			currentChunk = append(currentChunk, act)
		}
	}
	if len(currentChunk) > 0 {
		out = append(out, currentChunk)
	}
	return
}

func accountVariation(acct eos.AccountName, variation int) eos.AccountName {
	name := string(acct)
	if len(name) > 11 {
		name = name[:11]
	}
	variedName := name + string([]byte{'a' + byte(variation-1)})

	return eos.AccountName(variedName)
}
