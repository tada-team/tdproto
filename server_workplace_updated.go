package tdproto

func NewServerWorkplaceUpdated(workplaces []UserInfo) (r ServerWorkplaceUpdated) {
	r.Name = r.GetName()
	r.Params.Workplaces = workplaces
	return r
}

// ServerWorkplaceUpdated Workplace created or updated
type ServerWorkplaceUpdated struct {
	BaseEvent
	Params serverWorkplaceUpdatedParams `json:"params"`
}

func (p ServerWorkplaceUpdated) GetName() string { return "server.workplace.updated" }

// Params of the server.workplace.updated event
type serverWorkplaceUpdatedParams struct {
	// Workplaces info
	Workplaces []UserInfo `json:"workplaces"`
}
