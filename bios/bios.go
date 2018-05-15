package bios

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/crc64"
	"log"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/eoscanada/eos-bios/bios/disco"
	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
)

type BIOS struct {
	Network    *Network
	SingleOnly bool
	Log        *Logger

	LaunchDisco  *disco.Discovery
	TargetNetAPI *eos.API
	Snapshot     Snapshot
	BootSequence []*OperationType

	Genesis *GenesisJSON

	// ShuffledProducers is an ordered list of producers according to
	// the shuffled peers.
	ShuffledProducers []*Peer
	RandSource        rand.Source

	// MyPeers represent the peers my local node will handle. It is
	// plural because when launching a 3-node network, your peer will
	// be cloned a few time to have a full schedule of 21 producers.
	MyPeers []*Peer

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

	rawBootSeq, err := b.Network.ReadFromCache(bootseqFile)
	if err != nil {
		return fmt.Errorf("reading boot_sequence file: %s", err)
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
	fmt.Println("Starting Orchestraion process", time.Now())

	fmt.Println("Showing pre-randomized network discovered:")
	b.Network.PrintOrderedPeers()

	b.RandSource = b.waitLaunchBlock()

	// Once we have it, we can discover the net again (unless it's been discovered VERY recently)
	// and we b.Init() again.. so load the latest version of the LaunchData according to this
	// potentially new discovery network.
	fmt.Println("Seed network block used to seed randomization, updating graph one last time...")

	if err := b.Network.UpdateGraph(); err != nil {
		return fmt.Errorf("orchestrate: update graph: %s", err)
	}

	fmt.Println("Network used for launch:")
	b.Network.PrintOrderedPeers()

	if err := b.DispatchInit("orchestrate"); err != nil {
		return fmt.Errorf("dispatch init hook: %s", err)
	}

	switch b.MyRole() {
	case RoleBootNode:
		if err := b.RunBootSequence(); err != nil {
			return fmt.Errorf("orchestrate boot: %s", err)
		}
	case RoleABP:
		if err := b.RunJoinNetwork(true, true); err != nil {
			return fmt.Errorf("orchestrate join: %s", err)
		}
	default:
		if err := b.RunJoinNetwork(true, false); err != nil {
			return fmt.Errorf("orchestrate participate: %s", err)
		}
	}

	return b.DispatchDone("orchestrate")
}

func (b *BIOS) StartJoin(verify bool) error {
	fmt.Println("Starting network join process", time.Now())

	b.Network.PrintOrderedPeers()

	if err := b.DispatchInit("join"); err != nil {
		return fmt.Errorf("dispatch init hook: %s", err)
	}

	if err := b.RunJoinNetwork(verify, false); err != nil {
		return fmt.Errorf("boot network: %s", err)
	}

	return b.DispatchDone("join")
}

func (b *BIOS) StartBoot() error {
	fmt.Println("Starting network join process", time.Now())

	b.Network.PrintOrderedPeers()

	if err := b.DispatchInit("boot"); err != nil {
		return fmt.Errorf("dispatch init hook: %s", err)
	}

	if err := b.RunBootSequence(); err != nil {
		return fmt.Errorf("run boot sequence: %s", err)
	}

	return b.DispatchDone("boot")
}

func (b *BIOS) PrintOrderedPeers() {
	fmt.Println("###############################################################################################")
	fmt.Println("###################################  SHUFFLING RESULTS  #######################################")
	fmt.Println("")

	fmt.Printf("BIOS NODE: %s\n", b.ShuffledProducers[0].AccountName())
	for i := 1; i < 22 && len(b.ShuffledProducers) > i; i++ {
		fmt.Printf("ABP %02d:    %s\n", i, b.ShuffledProducers[i].AccountName())
	}
	fmt.Println("")
	fmt.Println("###############################################################################################")
	fmt.Println("########################################  BOOTING  ############################################")
	fmt.Println("")
	if b.AmIBootNode() {
		fmt.Println("I AM THE BOOT NODE! Let's get the ball rolling.")

	} else if b.AmIAppointedBlockProducer() {
		fmt.Println("I am NOT the BOOT NODE, but I AM ONE of the Appointed Block Producers. Stay tuned and watch the Boot node's media properties.")
	} else {
		fmt.Println("Okay... I'm not part of the Appointed Block Producers, we'll wait and be ready to join")
	}
	fmt.Println("")

	fmt.Println("###############################################################################################")
	fmt.Println("")
}

func (b *BIOS) RunBootSequence() error {
	fmt.Println("START BOOT SEQUENCE...")

	// keys, _ := b.TargetNetAPI.Signer.(*eos.KeyBag).AvailableKeys()
	// for _, key := range keys {
	// 	fmt.Println("Available key in the KeyBag:", key)
	// }

	ephemeralPrivateKey, err := b.GenerateEphemeralPrivKey()
	if err != nil {
		return err
	}

	b.EphemeralPrivateKey = ephemeralPrivateKey
	pubKey := ephemeralPrivateKey.PublicKey()
	b.EphemeralPublicKey = pubKey

	// b.TargetNetAPI.Debug = true

	privKey := ephemeralPrivateKey.String()

	fmt.Printf("Generated ephemeral keys: pub=%s priv=%s..%s\n", pubKey, privKey[:7], privKey[len(privKey)-7:])

	// Store keys in wallet, to sign `SetCode` and friends..
	if err := b.TargetNetAPI.Signer.ImportPrivateKey(privKey); err != nil {
		return fmt.Errorf("ImportWIF: %s", err)
	}

	genesisData := b.GenerateGenesisJSON(pubKey.String())

	if len(b.Network.MyPeer.Discovery.SeedNetworkPeers) > 0 && !b.SingleOnly {
		_, err := b.Network.SeedNetAPI.SignPushActions(
			disco.NewUpdateGenesis(b.Network.MyPeer.Discovery.SeedNetworkAccountName, genesisData, []string{}),
		)
		if err != nil {
			return fmt.Errorf("updating genesis on seednet: %s", err)
		}

		if err = b.DispatchBootPublishGenesis(genesisData); err != nil {
			return fmt.Errorf("dispatch boot_publish_genesis hook: %s", err)
		}
	}

	if err := b.DispatchBootNode(genesisData, pubKey.String(), privKey); err != nil {
		return fmt.Errorf("dispatch boot_node hook: %s", err)
	}

	fmt.Println(b.TargetNetAPI.Signer.AvailableKeys())

	for _, step := range b.BootSequence {
		fmt.Printf("%s  [%s]\n", step.Label, step.Op)

		acts, err := step.Data.Actions(b)
		if err != nil {
			return fmt.Errorf("getting actions for step %q: %s", step.Op, err)
		}

		if len(acts) != 0 {
			for idx, chunk := range chunkifyActions(acts, 400) { // transfers max out resources higher than ~400
				err := retry(5, 500*time.Millisecond, func() error {
					_, err = b.TargetNetAPI.SignPushActions(chunk...)
					if err != nil {
						return fmt.Errorf("SignPushActions for step %q, chunk %d: %s", step.Op, idx, err)
					}
					return nil
				})
				if err != nil {
					return err
				}
			}
		}
	}

	fmt.Println("Flushing transactions into blocks")
	time.Sleep(2 * time.Second)

	otherPeers := b.someTopmostPeersAddresses()
	if err := b.DispatchBootConnectMesh(otherPeers); err != nil {
		return fmt.Errorf("dispatch boot_connect_mesh: %s", err)
	}

	if err := b.DispatchBootPublishHandoff(); err != nil {
		return fmt.Errorf("dispatch boot_publish_handoff: %s", err)
	}

	return nil
}

type ActionMap map[string]*eos.Action

func (b *BIOS) RunJoinNetwork(verify, sabotage bool) error {

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

	if err := b.DispatchJoinNetwork(b.Genesis, b.MyPeers, otherPeers); err != nil {
		return fmt.Errorf("dispatch join_network hook: %s", err)
	}

	if verify {
		fmt.Println("###############################################################################################")
		fmt.Println("Launching chain verification")

		bootSeqActionMap := ActionMap{}
		bootSeqActionHexList := []string{}
		for _, step := range b.BootSequence {

			acts, err := step.Data.Actions(b)
			if err != nil {
				return fmt.Errorf("load boot seq, getting actions for step %q: %s", step.Op, err)
			}

			for _, stepAction := range acts {
				//fmt.Println("Verifying action type: ", reflect.TypeOf(stepAction.Data))
				data, err := eos.MarshalBinary(stepAction)
				if err != nil {
					return fmt.Errorf("load boot seq, marshalBinary action [%s], %s", stepAction.Name, err)
				}
				stepAction.SetToServer(false)
				key := hex.EncodeToString(data)

				//jsonData, err := json.Marshal(stepAction)
				//if err != nil {
				//	return fmt.Errorf("step action json marshalling, %s", err)
				//}
				//fmt.Println("Action Name: ", stepAction.Name)
				//fmt.Println("    json: ", string(jsonData))
				//fmt.Println("    hex data: ", key)
				//fmt.Println("    ----")

				bootSeqActionHexList = append(bootSeqActionHexList, key)
				if _, ok := bootSeqActionMap[key]; ok {
					log.Fatalf("Collision detected action [%s] with key [%s]\n", stepAction.Name, key)
				}
				bootSeqActionMap[key] = stepAction
			}
		}

		err := b.validateBootSeqActions(bootSeqActionMap, bootSeqActionHexList)
		if err != nil {
			return fmt.Errorf("- boot sequence verification endded with errors.\n%s", err)

		}

		fmt.Println("- Boot sequence verification passed")

		// Grab all the blocks from the chain
		// Compare each action, find it in our list
		// Use an ordered map ?

		fmt.Printf("- Verifying the `eosio` system account was properly disabled: ")
		for {
			time.Sleep(1 * time.Second)
			acct, err := b.TargetNetAPI.GetAccount(AN("eosio"))
			if err != nil {
				fmt.Printf("e")
				continue
			}

			if len(acct.Permissions) != 2 || acct.Permissions[0].RequiredAuth.Threshold != 0 || acct.Permissions[1].RequiredAuth.Threshold != 0 {
				// FIXME: perhaps check that there are no keys and
				// accounts.. that the account is *really* disabled.  we
				// can check elsewhere though.
				fmt.Printf(".")
				continue
			}

			fmt.Println(" OKAY")
			break
		}

		fmt.Println("Chain sync'd!")

		// IMPLEMENT THE BOOT SEQUENCE VERIFICATION.
		fmt.Println("")
		fmt.Println("All good! Chain verificaiton succeeded!")
		fmt.Println("")
	} else {
		fmt.Println("")
		fmt.Println("Not doing validation, the Appointed Block Producer will have done it.")
		fmt.Println("")
	}

	// TODO: loop operations, check all actions against blocks that you can fetch from here.
	// Check ALL actions, should match the orchestrated launch data:
	// - otherwise, sabotage

	fmt.Println("Awaiting for private key, for handoff verification.")
	fmt.Println("* This is the last step, and is done for the BIOS Boot node to prove it kept nothing to itself.")
	fmt.Println("")

	if verify {
		b.waitOnHandoff(b.Genesis)
	}

	return nil
}

func (b *BIOS) waitLaunchBlock() rand.Source {
	for {
		hash, err := b.Network.GetBlockHeight(b.LaunchDisco.SeedNetworkLaunchBlock)
		if err != nil {
			fmt.Println("couldn't fetch seed network block:", err)
		} else {

			if hash == "" {
				fmt.Println("block", b.LaunchDisco.SeedNetworkLaunchBlock, "not produced yet..")
			} else {
				bytes, err := hex.DecodeString(hash)
				if err != nil {
					fmt.Printf("block id is invalid hex: %q\n", hash)
				} else {
					chksum := crc64.Checksum(bytes, crc64.MakeTable(crc64.ECMA))
					return rand.NewSource(int64(chksum))
				}
			}
		}

		time.Sleep(2 * time.Second)
	}
}

func (b *BIOS) pollGenesisData() (genesis *GenesisJSON) {
	fmt.Println("")
	fmt.Println("Waiting for the BIOS Boot node to publish the genesis data to the seed network contract..")

	bootNode := b.ShuffledProducers[0]

	fmt.Printf("Polling...\n")
	for {
		time.Sleep(500 * time.Millisecond)

		genesisData, err := b.Network.PollGenesisTable(bootNode.Discovery.SeedNetworkAccountName)
		if err != nil {
			fmt.Printf("- data not ready: %s\n", err)
			continue
		}

		if len(genesisData) == 0 {
			fmt.Printf("- data still empty\n")
		}

		err = json.Unmarshal([]byte(genesisData), &genesis)
		if err != nil {
			fmt.Printf("- data not valid: %q (err=%s)\n", err, genesisData)
			continue
		}

		fmt.Println("Got genesis data:")
		fmt.Println("")
		fmt.Println("    ", genesisData)
		fmt.Println("")

		return
	}
}

func (b *BIOS) inputGenesisData() (genesis *GenesisJSON) {
	fmt.Println("")

	for {
		fmt.Printf("Please input the genesis data of the network you want to join: ")
		genesisData, err := ScanSingleLine()
		if err != nil {
			fmt.Println("error reading:", err)
			continue
		}

		err = json.Unmarshal([]byte(genesisData), &genesis)
		if err != nil {
			fmt.Printf("Invalid genesis data: %s\n", err)
			continue
		}

		return
	}
}

type ValidationError struct {
	BlockNumber       int
	Action            eos.Action
	Index             int
	ActionHexData     string
	PackedTransaction eos.PackedTransaction
}

func (e ValidationError) Error() string {
	s := fmt.Sprintf("Action [%d][%s] from block not found in boot sequences\n", e.Index, e.Action.Name)

	data, err := json.Marshal(e.Action)
	if err != nil {
		s += fmt.Sprintf("    json generation err : %s\n", err)
	} else {
		s += fmt.Sprintf("    json data : %s\n", string(data))
	}
	s += fmt.Sprintf("    hex data : %s\n", e.ActionHexData)
	//s += fmt.Sprintf("    packed transaction hex data : %s\n", hex.EncodeToString(e.PackedTransaction.PackedTransaction))

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

func (b *BIOS) validateBootSeqActions(bootSeqActionMap ActionMap, bootSeqActionHexList []string) (err error) {

	actionCount := 0
	blockCount := 1
	expectedActionCount := len(bootSeqActionHexList)
	validationErrors := make([]error, 0)
	//b.TargetNetAPI.Debug = true

	for actionCount < expectedActionCount {

		block, err := b.TargetNetAPI.GetBlockByNum(uint64(blockCount))
		if err != nil {
			return err
		}
		for _, blockTransaction := range block.Transactions {

			patchData := blockTransaction.Trx[1]
			var packedTransaction eos.PackedTransaction
			fmt.Println("Packed Trx: ", packedTransaction)
			err := json.Unmarshal(patchData, &packedTransaction)
			if err != nil {
				return fmt.Errorf("patching tx, %s", err)
			}
			signedTransaction, err := packedTransaction.Unpack()

			if err != nil {
				return fmt.Errorf("unpacking transaction, %s", err)
			}

			for _, action := range signedTransaction.Actions {

				data, err := eos.MarshalBinary(action)
				if err != nil {
					return fmt.Errorf("creating key with action, %s", err)
				}

				hexData := hex.EncodeToString(data)

				if _, ok := bootSeqActionMap[hexData]; !ok {

					validationErrors = append(validationErrors, ValidationError{
						BlockNumber:       blockCount,
						PackedTransaction: packedTransaction,
						Action:            *action,
						ActionHexData:     hexData,
						Index:             actionCount,
					})
				}

				actionCount++

				fmt.Printf("Verifing action [%s] [%d] of [%d] from block [%d]\n", action.Name, actionCount, expectedActionCount, blockCount)
			}
		}

		blockCount++

	}

	if len(validationErrors) > 0 {
		return ValidationErrors{Errors: validationErrors}
	}

	return nil
}

func (b *BIOS) waitOnHandoff(genesis *GenesisJSON) {
	for {
		fmt.Printf("Please paste the private key (or ipfs link): ")
		privKey, err := ScanSingleLine()
		if err != nil {
			fmt.Println("Error reading line:", err)
			continue
		}

		privKey = strings.TrimSpace(privKey)

		key, err := ecc.NewPrivateKey(privKey)
		if err != nil {
			fmt.Println("Invalid private key pasted:", err)
			continue
		}

		if key.PublicKey().String() == genesis.InitialKey {
			fmt.Println("")
			fmt.Println("   HANDOFF VERIFIED! EOS CHAIN IS ALIVE !")
			fmt.Println("")
			return
		} else {
			fmt.Println("")
			fmt.Println("   WARNING: private key provided does NOT match the genesis data")
			fmt.Println("")
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
		InitialChainID:   hex.EncodeToString(b.TargetNetAPI.ChainID),
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
	b.ShuffledProducers = b.Network.OrderedPeers()

	if b.RandSource != nil {
		b.shuffleProducers()
	}

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

				clonedProd := &Peer{
					ClonedAccountName: accountVariation(fromPeer.AccountName(), count),
					Discovery:         fromPeer.Discovery,
				}
				b.ShuffledProducers = append(b.ShuffledProducers, clonedProd)
			}
		}
	}

	return nil
}

func (b *BIOS) shuffleProducers() {
	fmt.Println("Shuffling producers listed in the launch file")
	r := rand.New(b.RandSource)
	// shuffle top 25%, capped to 5
	shuffleHowMany := int64(math.Min(math.Ceil(float64(len(b.ShuffledProducers))*0.25), 5))
	if shuffleHowMany > 1 {
		fmt.Println("- Shuffling top", shuffleHowMany)
		for round := 0; round < 100; round++ {
			from := r.Int63() % shuffleHowMany
			to := r.Int63() % shuffleHowMany
			if from == to {
				continue
			}

			//fmt.Println("Swapping from", from, "to", to)
			b.ShuffledProducers[from], b.ShuffledProducers[to] = b.ShuffledProducers[to], b.ShuffledProducers[from]
		}
	} else {
		fmt.Println("- No shuffling, network too small")
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

func chunkifyActions(actions []*eos.Action, chunkSize int) (out [][]*eos.Action) {
	currentChunk := []*eos.Action{}
	for _, act := range actions {
		if len(currentChunk) > chunkSize {
			out = append(out, currentChunk)
			currentChunk = []*eos.Action{}
		}
		currentChunk = append(currentChunk, act)
	}
	if len(currentChunk) > 0 {
		out = append(out, currentChunk)
	}
	return
}

func accountVariation(name string, variation int) string {
	if len(name) > 10 {
		name = name[:10]
	}
	return name + "." + string([]byte{'a' + byte(variation-1)})
}
