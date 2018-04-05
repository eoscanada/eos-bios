package main

type KickstartData struct {
	BIOSP2PAddress string `json:"bios_p2p_address"`
	PrivateKeyUsed string `json:"private_key_used"`
	PublicKeyUsed  string `json:"public_key_used"`
	GenesisJSON    string `json:"genesis_json"`
}
