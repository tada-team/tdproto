package tdproto

import (
	"time"
)

func NewServerCallBuzz(teamShort TeamShort, chatShort ChatShort, actorShort ContactShort, uid string, callType CallType, timeout time.Duration) (r ServerCallBuzz) {
	r.Name = r.GetName()
	r.Params.TeamShort = teamShort
	r.Params.ChatShort = chatShort
	r.Params.ActorShort = actorShort
	r.Params.Uid = uid
	r.Params.Jid = chatShort.Jid
	r.Params.BuzzTimeout = int(timeout.Seconds())
	r.Params.Icons = chatShort.Icons
	r.Params.DisplayName = chatShort.DisplayName
	r.Params.Team = teamShort.Uid
	r.Params.CallType = callType
	return r
}

// Call buzzing
type ServerCallBuzz struct {
	BaseEvent
	Params serverCallBuzzParams `json:"params"`
}

func (p ServerCallBuzz) GetName() string { return "server.call.buzz" }

// Params of the server.call.buzz event
type serverCallBuzzParams struct {
	// Short chat information
	ChatShort ChatShort `json:"chat"`

	// Short team information
	TeamShort TeamShort `json:"teaminfo"`

	// Chat icons
	Icons IconData `json:"icons"`

	// Chat or contact id
	Jid JID `json:"jid"`

	// Chat title
	DisplayName string `json:"display_name"`

	// Call id
	Uid string `json:"uid"`

	// Deprecated
	Team string `json:"team"`

	// CallType is a type of call("audio" - audio room, "video" - video room)
	CallType CallType `json:"call_type"`

	// Short call creator information
	ActorShort ContactShort `json:"actor"`

	// Number of seconds for stop buzzing
	BuzzTimeout int `json:"buzz_timeout"`
}
