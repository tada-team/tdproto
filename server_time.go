package tdproto

import "time"

func NewServerTime() (r ServerTime) {
	r.Name = r.GetName()
	r.Params.Time = IsoDatetime(time.Now())
	return r
}

type ServerTime struct {
	BaseEvent
	Params serverTimeParams `json:"params"`
}

func (p ServerTime) GetName() string { return "server.time" }

type serverTimeParams struct {
	Time string `json:"time"`
}
