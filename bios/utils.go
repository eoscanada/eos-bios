package bios

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	eos "github.com/eoscanada/eos-go"
)

func ScanLinesUntilBlank() (out string, err error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		var text string
		text, err = reader.ReadString('\n')
		//fmt.Println("Read line", text)
		if err != nil {
			return
		}

		out += text

		if text == "\n" {
			return strings.TrimSpace(out), nil
		}
	}
}

func ScanSingleLine() (out string, err error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}

// AN is a shortcut to create an AccountName
var AN = eos.AN

// PN is a shortcut to create a PermissionName
var PN = eos.PN

func flipEndianness(in uint64) (out uint64) {
	buf := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	binary.LittleEndian.PutUint64(buf, in)
	return binary.BigEndian.Uint64(buf)
}

func retry(attempts int, sleep time.Duration, callback func() error) (err error) {
	for i := 0; ; i++ {
		err = callback()
		if err == nil {
			return
		}

		if i >= (attempts - 1) {
			break
		}

		time.Sleep(sleep)

		log.Println("retrying after error:", err)
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

func AccountToNodeID(acct eos.AccountName) int64 {
	id, _ := eos.StringToName(string(acct))
	return int64(id)
}
