package bios

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
)

func PollEthereumClock(blockheight int) (blockhash string, err error) {
	// TODO: implement more sources, so we are distributed..  add the
	// possibility to paste the blockhash in case all sites are down?
	// do a direct connection to the Bitcoin blockchain using a swarm
	// of nodes ? That would be solid DDoS protection, but more
	// complex.

	useSite := rand.Int() % 2
	//useSite := 0
	switch useSite {
	case 0:
		return etherscanPollMethod(blockheight)
	case 1:
		return etherchainPollMethod(blockheight)
	}
	panic("hmm.. change the rand.Int modulo up here friend.. or the cases")
}

func etherscanPollMethod(height int) (blockhash string, err error) {
	var cnt string
	cnt, err = getURL(fmt.Sprintf("https://etherscan.io/block/%d", height))
	if err != nil {
		return
	}

	m := regexp.MustCompile(`>&nbsp;&nbsp;Hash:\n</td>\n<td>\n0x([0-9a-f]{64})\n`).FindStringSubmatch(cnt)
	if m == nil {
		// err = fmt.Errorf("regexp didn't match content on blockchain.info")
		return
	}

	return m[1], nil
}

func etherchainPollMethod(height int) (blockhash string, err error) {
	// TODO: Fix this
	var cnt string
	cnt, err = getURL(fmt.Sprintf("https://etherchain.io/block/%d", height))
	if err != nil {
		return
	}

	m := regexp.MustCompile(`>&nbsp;&nbsp;Hash:\n</td>\n<td>\n0x([0-9a-f]{64})\n`).FindStringSubmatch(cnt)
	if m == nil {
		// err = fmt.Errorf("regexp didn't match content on blockchain.info")
		return
	}

	return m[1], nil
}

func bitcoinPollMethodBlockchainInfo(height int) (blockhash string, err error) {
	var cnt string
	cnt, err = getURL(fmt.Sprintf("https://blockchain.info/block-height/%d", height))
	if err != nil {
		return
	}

	m := regexp.MustCompile(`<td>Hash</td>\n\s+<td><a href="/block/([0-9a-f]{64})">`).FindStringSubmatch(cnt)
	if m == nil {
		// err = fmt.Errorf("regexp didn't match content on blockchain.info")
		return
	}

	return m[1], nil
}

func bitcoinPollMethodBlockExplorer(height int) (blockhash string, err error) {
	var cnt string
	cnt, err = getURL(fmt.Sprintf("https://blockexplorer.com/api/block-index/%d", height))
	if err != nil {
		return
	}

	// {"status":"finished","blockChainHeight":520783,"syncPercentage":100,"height":520783,"error":null,"type":"bitcore node"}
	var out struct {
		Blockhash string
	}

	if err = json.Unmarshal([]byte(cnt), &out); err != nil {
		return "", nil
	}

	return out.Blockhash, nil

}

func getURL(u string) (string, error) {
	resp, err := http.Get(u)
	if err != nil {
		return "", err
	}

	cnt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(cnt), nil
}
