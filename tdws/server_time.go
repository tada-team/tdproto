package tdws

import "time"

func NewServerTime() (r ServerTime) {
	r.Name = r.GetName()
	r.Params.Time = IsoDatetime(time.Now())
	return r
}

// Current server time
type ServerTime struct {
	BaseEvent
	Params serverTimeParams `json:"params"`
}

func (p ServerTime) GetName() string { return "server.time" }

// Params of the server.time event
type serverTimeParams struct {
	// Current time
	Time ISODateTimeString `json:"time"`
}
