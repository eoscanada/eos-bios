package main

import (
	"bufio"
	"os"
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
