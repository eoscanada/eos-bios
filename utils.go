package main

import (
	"bufio"
	"os"

	eos "github.com/eosioca/eosapi"
)

func ScanLinesUntilBlank() (out string, err error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		var text string
		text, err = reader.ReadString('\n')
		if err != nil {
			return
		}

		out += text

		if text == "" {
			return
		}
	}
}

// AN is a shortcut to create an AccountName
var AN = eos.AN

// PN is a shortcut to create a PermissionName
var PN = eos.PN
