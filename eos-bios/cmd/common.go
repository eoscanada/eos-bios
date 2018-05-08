package cmd

import (
	"fmt"
	"os"

	"github.com/eoscanada/eos-bios"
	"github.com/eoscanada/eos-go"
	"github.com/ipfs/go-ipfs-api"
	"github.com/spf13/viper"
)

func fetchNetwork(single bool) (*bios.Network, error) {
	discoFile := viper.GetString("my-discovery")
	discovery, err := bios.LoadDiscoveryFromFile(discoFile)
	if err != nil {
		return nil, fmt.Errorf("loading %q: %s", discoFile, err)
	}

	ipfs := bios.NewIPFS(viper.GetString("ipfs-gateway-address"))

	seedNetAPI := eos.New(
		viper.GetString("seednet-api"),
		discovery.SeedNetworkChainID,
	)
	// FIXME: when the blockchain uses the chain ID, we'll set it (!!)
	seedNetAPI.ChainID = make([]byte, 32, 32)

	keyBag := eos.NewKeyBag()
	err = keyBag.ImportFromFile(viper.GetString("seednet-keys"))
	if err != nil {
		return nil, fmt.Errorf("importing keys: %s", err)
	}

	seedNetAPI.SetSigner(keyBag)

	net := bios.NewNetwork(
		viper.GetString("cache-path"),
		discovery,
		ipfs,
		seedNetworkContract, //	viper.GetString("seednet-contract"),
		seedNetAPI,
	)
	net.SingleOnly = single

	if err := net.UpdateGraph(); err != nil {
		return nil, fmt.Errorf("updating graph: %s", err)
	}

	return net, nil
}

func ipfsClient() (*shell.IdOutput, *shell.Shell) {
	ipfsClient := shell.NewShell(ipfsAPIAddress)

	fmt.Printf("Pinging ipfs node... ")
	info, err := ipfsClient.ID()
	if err != nil {
		fmt.Println("failed")
		fmt.Fprintf(os.Stderr, "error reaching ipfs api: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("ok")

	return info, ipfsClient
}

func setupBIOS(net *bios.Network) (b *bios.BIOS, err error) {
	targetNetAPI := eos.New(
		viper.GetString("target-api"),
		net.MyPeer.Discovery.TargetChainID,
	)
	// FIXME: this is until the blockchain (past dawn3) actually USES the chain ID
	// specified in the `genesis.json`.
	targetNetAPI.ChainID = make([]byte, 32, 32)

	keyBag := eos.NewKeyBag()
	targetNetAPI.SetSigner(keyBag)

	return bios.NewBIOS(net, targetNetAPI), nil
}
