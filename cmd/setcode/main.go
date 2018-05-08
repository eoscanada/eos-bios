package main

import (
	"bytes"
	"flag"
	"net/url"

	"log"

	"fmt"

	"path"

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
	"github.com/eoscanada/eos-go/system"
)

var accountNameString = flag.String("account-name", "shine", "Account name where to set shine code")
var codePath = flag.String("code-path", "../../eosio.disco", "Path to shine abi and wasm file")
var apiAddr = flag.String("api-addr", "http://localhost:8888", "RPC endpoint of the nodeos instance")
var privateKey = flag.String("private-key", "", "Private key")
var accountPublicKey = flag.String("account-public-key", "", "Shine account public key")

func main() {

	flag.Parse()

	apiAddrURL, err := url.Parse(*apiAddr)
	if err != nil {
		log.Fatalln("could not parse --api-addr:", err)
	}

	api := eos.New(apiAddrURL, bytes.Repeat([]byte{0}, 32))

	keyBag := eos.NewKeyBag()
	err = keyBag.Add(*privateKey)
	if err != nil {
		log.Fatal("key bag error: ", err)
	}

	api.SetSigner(keyBag)

	accountName := eos.AccountName(*accountNameString)
	accountResp, err := api.GetAccount(accountName)
	if err != nil {
		log.Fatal(err)
	}
	if len(accountResp.Permissions) == 0 {

		_, err = api.SignPushActions(
			system.NewNewAccount(eos.AccountName("eosio"), accountName, ecc.MustNewPublicKey(*accountPublicKey)),
		)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		println("User already existe ...")
	}

	setCodeTx, err := system.NewSetCodeTx(
		eos.AccountName(*accountNameString),
		path.Join(*codePath, "eosio.disco.wasm"),
		path.Join(*codePath, "eosio.disco.abi"),
	)
	if err != nil {
		log.Fatal("setcode creation error: ", err)
	}

	resp, err := api.SignPushTransaction(setCodeTx, &eos.TxOptions{})
	if err != nil {
		fmt.Println("ERROR calling set code:", err)
	} else {
		fmt.Println("All good: ", resp)
	}
}
