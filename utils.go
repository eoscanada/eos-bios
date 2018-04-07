package main

import (
	"bufio"
	"os"
	"strings"

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

// AN is a shortcut to create an AccountName
var AN = eos.AN

// PN is a shortcut to create a PermissionName
var PN = eos.PN
