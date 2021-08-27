package tdevents

import (
	"time"

	"github.com/tada-team/tdproto"
)

func NewServerCallBuzz(teamShort tdproto.TeamShort, chatShort tdproto.ChatShort, actorShort tdproto.ContactShort, uid string, timeout time.Duration) (r ServerCallBuzz) {
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
	// Chat or contact id
	Jid tdproto.JID `json:"jid"`

	// Chat icons
	Icons tdproto.IconData `json:"icons"`

	// Chat title
	DisplayName string `json:"display_name"`

	// Short team information
	TeamShort tdproto.TeamShort `json:"teaminfo"`

	// Short chat information
	ChatShort tdproto.ChatShort `json:"chat"`

	// Short call creator information
	ActorShort tdproto.ContactShort `json:"actor"`

	// Call id
	Uid string `json:"uid"`

	// Number of seconds for stop buzzing
	BuzzTimeout int `json:"buzz_timeout"`

	// Deprecated
	Team string `json:"team"`
}
