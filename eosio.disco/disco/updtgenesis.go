package disco

import "github.com/eoscanada/eos-go"

func NewUpdateGenesis(account eos.AccountName, genesisJSON string, initialP2PAddresses []string) *eos.Action {

	action := &eos.Action{
		Account: eos.AccountName("eosio.disco"),
		Name:    eos.ActionName("updtgenesis"),
		Authorization: []eos.PermissionLevel{
			{account, eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(UpdtGenesis{
			Account:             account,
			GenesisJSON:         genesisJSON,
			InitialP2PAddresses: initialP2PAddresses,
		}),
	}
	return action
}

type UpdtGenesis struct {
	Account             eos.AccountName `json:"account"`
	GenesisJSON         string          `json:"genesis_json"`
	InitialP2PAddresses []string        `json:"initial_p2p_addresses"`
}
