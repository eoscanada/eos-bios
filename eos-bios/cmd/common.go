package cmd

import (
	"fmt"
	"os"

	"github.com/eoscanada/eos-bios"
	"github.com/eoscanada/eos-go"
	"github.com/ipfs/go-ipfs-api"
	"github.com/spf13/viper"
)

func fetchNetwork(api *eos.API, ipfs *bios.IPFS, contractName string) (*bios.Network, error) {

	net := bios.NewNetwork(cachePath, myDiscoveryFile, ipfs, contractName, api)

	net.UseCache = viper.GetBool("no-discovery")

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
