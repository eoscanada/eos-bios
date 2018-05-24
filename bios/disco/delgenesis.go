package disco

import "github.com/eoscanada/eos-go"

func NewDeleteGenesis(account eos.AccountName) *eos.Action {
	action := &eos.Action{
		Account: eos.AccountName("eosio.disco"),
		Name:    eos.ActionName("delgenesis"),
		Authorization: []eos.PermissionLevel{
			{account, eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(DelGenesis{
			Account: account,
		}),
	}
	return action
}

type DelGenesis struct {
	Account eos.AccountName `json:"account"`
}
