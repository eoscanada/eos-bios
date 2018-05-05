package bios

import (
	"encoding/json"
	"io/ioutil"

	yaml2json "github.com/bronze1man/go-yaml2json"
)

func yamlUnmarshal(cnt []byte, v interface{}) error {
	jsonCnt, err := yaml2json.Convert(cnt)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonCnt, v)
}

func ValidateDiscoveryFile(filename string) error {
	cnt, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var disco *Discovery
	err = yamlUnmarshal(cnt, &disco)
	if err != nil {
		return err
	}

	return nil
}
