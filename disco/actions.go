package disco

import "github.com/eoscanada/eos-go"

func NewUpdateDiscovery(discovery *Discovery, account eos.AccountName) *eos.Action {

	action := &eos.Action{
		Account: eos.AccountName("eosio"),
		Name:    eos.ActionName("discovery"),
		Authorization: []eos.PermissionLevel{
			{account, eos.PermissionName("active")},
		},
		Data: eos.NewActionData(discovery),
	}
	return action
}
