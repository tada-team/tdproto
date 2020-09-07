package tdproto

import "time"

func NewServerTime() (r ServerTime) {
	r.BaseEvent.Name = "server.time"
	r.Params.Time = IsoDatetime(time.Now())
	return r
}

type ServerTime struct {
	BaseEvent
	Params ServerTimeParams `json:"params"`
}

type ServerTimeParams struct {
	Time string `json:"time"`
}
