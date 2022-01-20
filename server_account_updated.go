package tdproto

func NewServerAccountUpdated(account PersonalAccountBilling) (r ServerAccountUpdated) {
	r.Name = r.GetName()
	r.Params.Account = account
	return r
}

// Personal Account created or updated
type ServerAccountUpdated struct {
	BaseEvent
	Params serverAccountUpdatedParams `json:"params"`
}

func (p ServerAccountUpdated) GetName() string { return "server.account.updated" }

// Params of the server.account.updated event
type serverAccountUpdatedParams struct {
	// Personal Account info
	Account PersonalAccountBilling `json:"account"`
}
