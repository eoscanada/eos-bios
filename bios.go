package bios

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
)

type BIOS struct {
	Network *Network

	LaunchData   *LaunchData
	EOSAPI       *eos.API
	Snapshot     Snapshot
	BootSequence []*OperationType

	KickstartData *KickstartData

	// ShuffledProducers is an ordered list of producers according to
	// the shuffled peers.
	ShuffledProducers []*Peer

	// MyPeers represent the peers my local node will handle. It is
	// plural because when launching a 3-node network, your peer will
	// be cloned a few time to have a full schedule of 21 producers.
	MyPeers []*Peer

	EphemeralPrivateKey *ecc.PrivateKey
}

func NewBIOS(network *Network, api *eos.API) *BIOS {
	b := &BIOS{
		Network: network,
		EOSAPI:  api,
	}
	return b
}

func (b *BIOS) SetKickstartData(ks *KickstartData) {
	b.KickstartData = ks
}

func (b *BIOS) Init() error {
	// Load launch data
	launchData, err := b.Network.ConsensusLaunchData()
	if err != nil {
		fmt.Println("couldn'get consensus on launch data:", err)
		os.Exit(1)
	}

	b.LaunchData = launchData

	// TODO: check that nodes that are ABP or participants do have an
	// EOSIOABPSigningKey set.

	// Load the boot sequence
	rawBootSeq, err := b.Network.ReadFromCache(launchData.BootSequence)
	if err != nil {
		return fmt.Errorf("reading boot_sequence file: %s", err)
	}

	var bootSeq struct {
		BootSequence []*OperationType `json:"boot_sequence"`
	}
	if err := yamlUnmarshal(rawBootSeq, &bootSeq); err != nil {
		return fmt.Errorf("loading boot sequence: %s", err)
	}

	b.BootSequence = bootSeq.BootSequence

	// Load snapshot data
	if launchData.Snapshot != "" {
		rawSnapshot, err := b.Network.ReadFromCache(launchData.Snapshot)
		if err != nil {
			return fmt.Errorf("reading snapshot file: %s", err)
		}

		snapshotData, err := NewSnapshot(rawSnapshot)
		if err != nil {
			return fmt.Errorf("loading snapshot csv: %s", err)
		}
		b.Snapshot = snapshotData
	}

	if err := b.shuffleProducers(); err != nil {
		return err
	}

	if err = b.setMyPeers(); err != nil {
		return fmt.Errorf("error setting my producer definitions: %s", err)
	}

	return nil
}

func (b *BIOS) StartOrchestrate(secretP2PAddress string) error {
	fmt.Println("Starting Orchestraion process", time.Now())

	b.Network.PrintOrderedPeers()

	if err := b.DispatchInit(); err != nil {
		return fmt.Errorf("failed init hook: %s", err)
	}

	switch b.MyRole() {
	case RoleBootNode:
		if err := b.RunBootSequence(secretP2PAddress); err != nil {
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

	return b.DispatchDone()
}

func (b *BIOS) StartJoin(verify bool) error {
	fmt.Println("Starting network join process", time.Now())

	b.Network.PrintOrderedPeers()

	if err := b.DispatchInit(); err != nil {
		return fmt.Errorf("failed init hook: %s", err)
	}

	if err := b.RunJoinNetwork(verify, false); err != nil {
		return fmt.Errorf("boot network: %s", err)
	}

	return b.DispatchDone()
}

func (b *BIOS) StartBoot(secretP2PAddress string) error {
	fmt.Println("Starting network join process", time.Now())

	b.Network.PrintOrderedPeers()

	if err := b.DispatchInit(); err != nil {
		return fmt.Errorf("failed init hook: %s", err)
	}

	if err := b.RunBootSequence(secretP2PAddress); err != nil {
		return fmt.Errorf("join network: %s", err)
	}

	return b.DispatchDone()
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

func (b *BIOS) RunBootSequence(secretP2PAddress string) error {
	fmt.Println("START BOOT SEQUENCE...")

	ephemeralPrivateKey, err := b.GenerateEphemeralPrivKey()
	if err != nil {
		return err
	}

	b.EphemeralPrivateKey = ephemeralPrivateKey

	// b.EOSAPI.Debug = true

	pubKey := ephemeralPrivateKey.PublicKey().String()
	privKey := ephemeralPrivateKey.String()

	fmt.Printf("Generated ephemeral keys: pub=%s priv=%s..%s\n", pubKey, privKey[:7], privKey[len(privKey)-7:])

	// Store keys in wallet, to sign `SetCode` and friends..
	if err := b.EOSAPI.Signer.ImportPrivateKey(privKey); err != nil {
		return fmt.Errorf("ImportWIF: %s", err)
	}

	keys, _ := b.EOSAPI.Signer.(*eos.KeyBag).AvailableKeys()
	for _, key := range keys {
		fmt.Println("Available key in the KeyBag:", key)
	}

	genesisData := b.GenerateGenesisJSON(pubKey)

	if err = b.DispatchBootNetwork(genesisData, pubKey, privKey); err != nil {
		return fmt.Errorf("dispatch config_ready hook: %s", err)
	}

	fmt.Println(b.EOSAPI.Signer.AvailableKeys())

	for _, step := range b.BootSequence {
		fmt.Printf("%s  [%s]\n", step.Label, step.Op)

		acts, err := step.Data.Actions(b)
		if err != nil {
			return fmt.Errorf("getting actions for step %q: %s", step.Op, err)
		}

		if len(acts) != 0 {
			for idx, chunk := range chunkifyActions(acts, 400) { // transfers max out resources higher than ~400
				_, err = b.EOSAPI.SignPushActions(chunk...)
				if err != nil {
					return fmt.Errorf("SignPushActions for step %q, chunk %d: %s", step.Op, idx, err)
				}
			}
		}
	}

	fmt.Println("Preparing kickstart data")

	kickstartData := &KickstartData{
		BIOSP2PAddress: secretP2PAddress,
		PublicKeyUsed:  pubKey,
		PrivateKeyUsed: privKey,
		GenesisJSON:    genesisData,
	}
	kd, _ := json.Marshal(kickstartData)
	ksdata := base64.RawStdEncoding.EncodeToString(kd)

	// TODO: encrypt it for those who need it

	fmt.Println("PUBLISH THIS KICKSTART DATA:")
	fmt.Println("")
	fmt.Println(ksdata)
	fmt.Println("")

	if err = b.DispatchPublishKickstartData(ksdata); err != nil {
		return fmt.Errorf("dispatch publish_kickstart_data: %s", err)
	}

	return nil
}

func (b *BIOS) RunJoinNetwork(verify, sabotage bool) error {
	if b.KickstartData == nil {
		kickstart, err := b.waitOnKickstartData()
		if err != nil {
			return err
		}
		b.KickstartData = &kickstart
	}

	if err := b.DispatchJoinNetwork(b.KickstartData, b.MyPeers); err != nil {
		return err
	}

	fmt.Println("Not doing any validation, the ABPs have done it")

	if verify {
		fmt.Println("###############################################################################################")
		fmt.Println("Launching chain verification")

		// Grab all the Actions, serialize them.
		// Grab all the blocks from the chain
		// Compare each action, find it in our list
		// Use an ordered map ?
		// for _, step := range b.BootSequence {
		// 	fmt.Printf("%s  [%s]\n", step.Label, step.Op)

		// 	acts, err := step.Data.Actions(b)
		// 	if err != nil {
		// 		return fmt.Errorf("getting actions for step %q: %s", step.Op, err)
		// 	}

		// }

		fmt.Printf("- Verifying the `eosio` system account was properly disabled: ")
		for {
			time.Sleep(1 * time.Second)
			acct, err := b.EOSAPI.GetAccount(AN("eosio"))
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
	}

	// TODO: loop operations, check all actions against blocks that you can fetch from here.
	// Do all the checks:
	//  - all Producers are properly setup
	//  - anything fails, SABOTAGE
	// Publish a PGP Signed message with your local IP.. push to properties
	// Dispatch webhook PublishKickstartPublic (with a Kickstart Data object)

	return nil
}

func (b *BIOS) waitOnKickstartData() (kickstart KickstartData, err error) {
	fmt.Println("")
	fmt.Println("The BIOS node will publish the Kickstart Data through their social media.")
	bootNode := b.ShuffledProducers[0]
	disco := bootNode.Discovery
	if disco.Website != "" {
		fmt.Println("  Main website:", disco.Website)
	}
	if disco.SocialTwitter != "" {
		fmt.Println("  Twitter:", disco.SocialTwitter)
	}
	if disco.SocialFacebook != "" {
		fmt.Println("  Facebook:", disco.SocialFacebook)
	}
	if disco.SocialTelegram != "" {
		fmt.Println("  Telegram:", disco.SocialTelegram)
	}
	if disco.SocialSlack != "" {
		fmt.Println("  Slack:", disco.SocialSlack)
	}
	if disco.SocialSteem != "" {
		fmt.Println("  Steem:", disco.SocialSteem)
	}
	if disco.SocialSteemIt != "" {
		fmt.Println("  SteemIt:", disco.SocialSteemIt)
	}
	if disco.SocialKeybase != "" {
		fmt.Println("  Keybase:", disco.SocialKeybase)
	}
	if disco.SocialWeChat != "" {
		fmt.Println("  WeChat:", disco.SocialWeChat)
	}
	if disco.SocialYouTube != "" {
		fmt.Println("  YouTube:", disco.SocialYouTube)
	}
	if disco.SocialGitHub != "" {
		fmt.Println("  GitHub:", disco.SocialGitHub)
	}
	// TODO: print the social media properties of the BP..
	fmt.Println("Paste it here and finish with two blank lines (ENTER twice):")
	fmt.Println("")

	// Wait on stdin for kickstart data (will we have some other polling / subscription mechanisms?)
	//    Accept any base64, unpadded, multi-line until we receive a blank line, concat and decode.
	// FIXME: this is a quick hack to just pass the p2p address
	lines, err := ScanLinesUntilBlank()
	if err != nil {
		return
	}

	rawKickstartData, err := base64.RawStdEncoding.DecodeString(strings.Replace(strings.TrimSpace(lines), "\n", "", -1))
	if err != nil {
		return kickstart, fmt.Errorf("kickstart base64 decode: %s", err)
	}

	err = json.Unmarshal(rawKickstartData, &kickstart)
	if err != nil {
		return kickstart, fmt.Errorf("unmarshal kickstart data: %s", err)
	}

	privKey, err := ecc.NewPrivateKey(kickstart.PrivateKeyUsed)
	if err != nil {
		return kickstart, fmt.Errorf("unable to load private key %q: %s", kickstart.PrivateKeyUsed, err)
	}

	b.EphemeralPrivateKey = privKey

	// TODO: check if the privKey corresponds to the public key sent, if not, we should
	// drop that kickstart data.. and listen to another one..

	return
}

func (b *BIOS) GenerateEphemeralPrivKey() (*ecc.PrivateKey, error) {
	return ecc.NewRandomPrivateKey()
}

func (b *BIOS) GenerateGenesisJSON(pubKey string) string {
	// known not to fail
	cnt, _ := json.Marshal(&GenesisJSON{
		InitialTimestamp: time.Now().UTC().Format("2006-01-02T15:04:05"), // TODO: becomes the bitcoin block if we use that as a seed for randomization? just the current time/date ?
		InitialKey:       pubKey,
		InitialChainID:   hex.EncodeToString(b.EOSAPI.ChainID),
	})
	return string(cnt)
}

func (b *BIOS) shuffleProducers() error {
	fmt.Println("Shuffling producers listed in the launch file [NOT IMPLEMENTED]")

	b.ShuffledProducers = b.Network.OrderedPeers()

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

func (b *BIOS) IsBootNode(account string) bool {
	return string(b.ShuffledProducers[0].AccountName()) == account
}

func (b *BIOS) AmIBootNode() bool {
	return b.IsBootNode(b.Network.MyPeer.Discovery.EOSIOAccountName)
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
		if b.ShuffledProducers[i].Discovery.EOSIOAccountName == account {
			return true
		}
	}
	return false
}

func (b *BIOS) AmIAppointedBlockProducer() bool {
	return b.IsAppointedBlockProducer(b.Network.MyPeer.Discovery.EOSIOAccountName)
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
		if peer.Discovery.EOSIOAccountName == myPeer.Discovery.EOSIOAccountName {
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
