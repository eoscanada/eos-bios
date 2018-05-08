package cmd

import (
	"fmt"
	"os"

	"github.com/eoscanada/eos-bios"
	"github.com/eoscanada/eos-go"
	"github.com/ipfs/go-ipfs-api"
	"github.com/spf13/viper"
)

func fetchNetwork(api *eos.API) (*bios.Network, error) {
	discovery, err := bios.LoadDiscoveryFromFile(viper.GetString("my-discovery"))
	if err != nil {
		return nil, err
	}

	ipfs := bios.NewIPFS(viper.GetString("ipfs-gateway-address"))

	seedNetAPI := eos.New(
		viper.GetString("seednet-api"),
		discovery.SeedNetworkChainID,
	)

	net := bios.NewNetwork(
		viper.GetString("cache-path"),
		discovery,
		ipfs,
		viper.GetString("seednet-contract"),
		seedNetAPI,
	)

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

func api(chainID eos.SHA256Bytes) (api *eos.API, err error) {
	api = eos.New(apiAddressURL, chainID)
	keyBag := eos.NewKeyBag()
	err = keyBag.ImportFromFile(seedNetworkKeysFile)
	return
}
