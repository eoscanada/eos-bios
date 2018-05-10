package disco

import "github.com/eoscanada/eos-go"

func NewUpdateDiscovery(account eos.AccountName, discovery *Discovery) *eos.Action {
	action := &eos.Action{
		Account: eos.AccountName("eosio.disco"),
		Name:    eos.ActionName("updtdisco"),
		Authorization: []eos.PermissionLevel{
			{account, eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(UpdtDisco{
			Account:   account,
			Discovery: discovery,
		}),
	}
	return action
}

type UpdtDisco struct {
	Account   eos.AccountName `json:"account"`
	Discovery *Discovery      `json:"disco"`
}
