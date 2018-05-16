package bios

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/crc64"
	"io/ioutil"
	"log"
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
	// * a second that can be recalled just after receiving the Launch Block, which
	//   might change the bootsequence we've agreed upon, might change the network,
	//   topology, etc..

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
		fmt.Printf("Using overridden boot sequence from %q\n", b.OverrideBootSequenceFile)

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
	fmt.Println("Starting Orchestraion process", time.Now())

	fmt.Println("Showing pre-randomized network discovered:")
	b.PrintProducerSchedule()

	b.RandSource = b.waitLaunchBlock()

	// Once we have it, we can discover the net again (unless it's been discovered VERY recently)
	// and we b.Init() again.. so load the latest version of the LaunchData according to this
	// potentially new discovery network.
	fmt.Println("Seed network block used to seed randomization, updating graph one last time...")

	if err := b.Network.UpdateGraph(); err != nil {
		return fmt.Errorf("update graph: %s", err)
	}

	fmt.Println("Network used for launch:")
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

func (b *BIOS) StartJoin(verify bool) error {
	fmt.Println("Starting network join process", time.Now())

	b.PrintProducerSchedule()

	if err := b.DispatchInit("join"); err != nil {
		return fmt.Errorf("dispatch init hook: %s", err)
	}

	if err := b.RunJoinNetwork(verify, false); err != nil {
		return fmt.Errorf("join network: %s", err)
	}

	return b.DispatchDone("join")
}

func (b *BIOS) StartBoot() error {
	fmt.Println("Starting network join process", time.Now())

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

	fmt.Println("")
	fmt.Println("###############################################################################################")
	fmt.Println("")
	if b.AmIBootNode() {
		fmt.Println("                              MY ROLE: BIOS BOOT NODE")
	} else if b.AmIAppointedBlockProducer() {
		fmt.Println("                              MY ROLE: APPOINTED BLOCK PRODUCER")
	} else {
		fmt.Println("                              MY ROLE: JOINING NETWORK")
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

	fmt.Printf("Generated ephemeral keys:\n\n\tPublic key: %s\n\tPrivate key: %s..%s\n\n", pubKey, privKey[:7], privKey[len(privKey)-7:])

	// Store keys in wallet, to sign `SetCode` and friends..
	if err := b.TargetNetAPI.Signer.ImportPrivateKey(privKey); err != nil {
		return fmt.Errorf("ImportWIF: %s", err)
	}

	genesisData := b.GenerateGenesisJSON(pubKey.String())

	if len(b.Network.MyPeer.Discovery.SeedNetworkPeers) > 0 && !b.SingleOnly {

		fmt.Printf("Publishing genesis data to the seed network... ")
		_, err := b.Network.SeedNetAPI.SignPushActions(
			disco.NewUpdateGenesis(b.Network.MyPeer.Discovery.SeedNetworkAccountName, genesisData, []string{}),
		)
		if err != nil {
			fmt.Println("")
			return fmt.Errorf("updating genesis on seednet: %s", err)
		}
		fmt.Println(" done")

		if err = b.DispatchBootPublishGenesis(genesisData); err != nil {
			return fmt.Errorf("dispatch boot_publish_genesis hook: %s", err)
		}
	}

	if err := b.DispatchBootNode(genesisData, pubKey.String(), privKey); err != nil {
		return fmt.Errorf("dispatch boot_node hook: %s", err)
	}

	fmt.Println("In-memory keys:")
	fmt.Println(b.TargetNetAPI.Signer.AvailableKeys())
	fmt.Println("")

	// eos.Debug = true

	for _, step := range b.BootSequence {
		fmt.Printf("%s  [%s] ", step.Label, step.Op)

		acts, err := step.Data.Actions(b)
		if err != nil {
			return fmt.Errorf("getting actions for step %q: %s", step.Op, err)
		}

		if len(acts) != 0 {
			for idx, chunk := range chunkifyActions(acts, 250) { // transfers max out resources higher than ~400
				err := retry(5, 500*time.Millisecond, func() error {
					_, err := b.TargetNetAPI.SignPushActions(chunk...)
					if err != nil {
						if strings.Contains(err.Error(), `"message":"itr != structs.end(): Unknown struct ","file":"abi_serializer.cpp"`) { // server-side error for serializing, but the transaction went through !!
							return nil
						}
						return fmt.Errorf("SignPushActions for step %q, chunk %d: %s", step.Op, idx, err)
					}
					return nil
				})
				if err != nil {
					fmt.Printf(" error\n")
					return err
				}
				fmt.Printf(" done\n")
			}
		}
	}

	fmt.Println("Flushing transactions into blocks")
	time.Sleep(2 * time.Second)

	orderedPeers := b.Network.OrderedPeers(b.Network.MyNetwork())

	otherPeers := b.someTopmostPeersAddresses(orderedPeers)
	if err := b.DispatchBootConnectMesh(otherPeers); err != nil {
		return fmt.Errorf("dispatch boot_connect_mesh: %s", err)
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

	if err := b.DispatchJoinNetwork(b.Genesis, b.getMyPeerVariations(), otherPeers); err != nil {
		return fmt.Errorf("dispatch join_network hook: %s", err)
	}

	if verify {
		fmt.Println("###############################################################################################")
		fmt.Println("Launching chain verification")
		fmt.Println("")
		fmt.Println("  - VALIDATION IS BEING IMPLEMENTED BY CHARLES ! Give him a day or so ! :)")
		fmt.Println("")
		fmt.Println("DONE")
		os.Exit(0)

		seedActionMap, err := b.fetchActions()
		if err != nil {
			return fmt.Errorf("verifing, fetching actions from seed, %s", err)
		}

		fmt.Println("Seed actions found: ", len(seedActionMap))

		for _, step := range b.BootSequence {
			fmt.Printf("%s  [%s]\n", step.Label, step.Op)

			acts, err := step.Data.Actions(b)
			if err != nil {
				return fmt.Errorf("verifing, getting actions for step %q: %s", step.Op, err)
			}

			for _, stepAction := range acts {
				//fmt.Println("Verifying action type: ", reflect.TypeOf(stepAction.Data))
				data, err := eos.MarshalBinary(stepAction.Data)
				if err != nil {
					return fmt.Errorf("verifying, marshalBinary, %s", err)
				}
				key := hex.EncodeToString(data)
				//fmt.Println("verifying key : ", key)

				if seedAction, ok := seedActionMap[key]; !ok {
					//return fmt.Errorf("verify, step action [%s] does not validate", stepAction.Name)
					fmt.Printf("✘ verify, step action [%s] does not validate\n", stepAction.Name)
				} else {
					fmt.Printf("✔ action [%s] verified\n", seedAction[0].Name)
				}
			}

		}

		//***********************************************************************
		//***********************************************************************
		log.Fatal("let's crash!")
		//***********************************************************************
		//***********************************************************************

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
	targetBlockNum := uint32(b.LaunchDisco.SeedNetworkLaunchBlock)

	fmt.Println("Polling seed network until launch block, target:", targetBlockNum)

	for {
		lastBlockNum, err := b.Network.GetLastBlockNum()
		if err != nil {
			fmt.Println("error fetching seed network's latest block num:", err)
		}

		if lastBlockNum < targetBlockNum {
			fmt.Printf("- not yet, %d seconds to go\n", (targetBlockNum-lastBlockNum)/2)
			time.Sleep(time.Second)
			continue
		}

		// GET INFO and check if the block is the right height, otherwise, give a prediction of
		// when it will arrive :) countdown!
		hash, err := b.Network.GetBlockHeight(targetBlockNum)
		if err != nil {
			fmt.Println("error fetching seed network's target block height:", err)
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Println("- got block", targetBlockNum, "- hash is", hash)
		bytes, _ := hex.DecodeString(hash)
		chksum := crc64.Checksum(bytes, crc64.MakeTable(crc64.ECMA))
		return rand.NewSource(int64(chksum))

	}
}

func (b *BIOS) pollGenesisData() (genesis *GenesisJSON) {
	fmt.Println("")
	fmt.Println("Waiting for the BIOS Boot node to publish the genesis data to the seed network contract..")

	bootNode := b.ShuffledProducers[0]

	fmt.Printf("Polling..")
	for {
		time.Sleep(500 * time.Millisecond)

		fmt.Printf(".")
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

		fmt.Println("")
		fmt.Println("Got genesis data:")
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

type ActionMap map[string][]*eos.Action

func (b *BIOS) fetchActions() (actions ActionMap, err error) {
	accounts := []eos.AccountName{
		eos.AccountName("eosio"),
		eos.AccountName("eosio.msig"),
		eos.AccountName("eosio.disco"),
		eos.AccountName("eosio.token"),
	}
	actions = ActionMap{}
	for _, account := range accounts {
		fmt.Printf("Fecthing actions for account [%s]\n", account)
		out, err := b.Network.SeedNetAPI.GetTransactions(account)
		if err != nil {
			err = fmt.Errorf("fectching transactions for [%s] account, %s", account, err)
		}

		for _, tx := range out.Transactions {
			for _, action := range tx.Transaction.Transaction.Actions {

				key := hex.EncodeToString(action.HexData)
				fmt.Printf("action [%s]\n", action.Name)
				//fmt.Printf("action [%s] key : %s\n", action.Name, key)
				//data, err := json.Marshal(action)
				//assert.NoError(t, err)
				//fmt.Println("Data  : ", string(data))
				//if collision, ok := actions[key]; ok {
				//	cdata, err := json.Marshal(collision)
				//	assert.NoError(t, err)
				//	fmt.Println("CData : ", string(cdata))
				//	fmt.Println("Found a colision")
				//}

				actions[key] = append(actions[key], action)
			}
		}
	}
	return
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
				clonedDisco.TargetAccountName = accountVariation(fromPeer.Discovery.TargetAccountName, count)
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

func (b *BIOS) shuffleProducers() {
	if b.RandSource == nil {
		fmt.Println("Random source not set, skipping producer shuffling")
		return
	}

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

func accountVariation(acct eos.AccountName, variation int) eos.AccountName {
	name := string(acct)
	if len(name) > 10 {
		name = name[:10]
	}
	variedName := name + "." + string([]byte{'a' + byte(variation-1)})

	return eos.AccountName(variedName)
}
