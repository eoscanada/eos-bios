package cmd

import (
	"fmt"
	"os"

	"github.com/eoscanada/eos-bios/bios"
	"github.com/eoscanada/eos-go"
	"github.com/ipfs/go-ipfs-api"
	"github.com/spf13/viper"
)

func fetchNetwork(single, downloadRefs bool) (*bios.Network, error) {
	discoFile := viper.GetString("my-discovery")
	discovery, err := bios.LoadDiscoveryFromFile(discoFile)
	if err != nil {
		return nil, fmt.Errorf("loading %q: %s", discoFile, err)
	}

	ipfs := bios.NewIPFS(viper.GetString("ipfs"))

	seedNetHTTP := viper.GetString("seednet-api")
	if seedNetHTTP == "" {
		seedNetHTTP = discovery.SeedNetworkHTTPAddress
	}
	if seedNetHTTP == "" && !single {
		return nil, fmt.Errorf("missing `seed_network_http_address` and no `--seednet-api` override provided")
	}

	seedNetAPI := eos.New(seedNetHTTP)

	keyBag := eos.NewKeyBag()
	err = keyBag.ImportFromFile(viper.GetString("seednet-keys"))
	if err != nil {
		fmt.Println("WARN: you might want to simply rename privkeys.keys to seed_network.keys")
		return nil, fmt.Errorf("importing keys: %s", err)
	}

	seedNetAPI.SetSigner(keyBag)

	logger := bios.NewLogger()
	logger.Debug = viper.GetBool("verbose")

	net := bios.NewNetwork(
		viper.GetString("cache-path"),
		discovery,
		ipfs,
		seedNetworkContract, //	viper.GetString("seednet-contract"),
		seedNetAPI,
	)
	net.Log = logger

	if single {
		net.SetLocalNetwork()
	}

	if err := net.UpdateGraph(); err != nil {
		return nil, fmt.Errorf("updating graph: %s", err)
	}

	if downloadRefs {
		if err := net.DownloadReferences(); err != nil {
			return nil, fmt.Errorf("downloading content: %s", err)
		}
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
	targetNetHTTP := viper.GetString("target-api")
	if targetNetHTTP == "" {
		targetNetHTTP = net.MyPeer.Discovery.TargetHTTPAddress
	}
	if targetNetHTTP == "" {
		return nil, fmt.Errorf("missing `target_http_address` and no `--target-api` override provided")
	}

	targetNetAPI := eos.New(targetNetHTTP)
	targetNetAPI.SetSigner(eos.NewKeyBag())

	if viper.GetBool("fast-inject") {
		targetNetAPI.EnableKeepAlives()
	}

	b = bios.NewBIOS(net.Log, net, targetNetAPI)
	b.WriteActions = viper.GetBool("write-actions")
	b.HackVotingAccounts = viper.GetBool("hack-voting-accounts")
	return b, nil
}
