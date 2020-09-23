package tdproto

import (
	"time"
)

func NewServerCallBuzz(teamShort TeamShort, chatShort ChatShort, actorShort ContactShort, uid string, timeout time.Duration) (r ServerCallBuzz) {
	r.BaseEvent.Name = "server.call.buzz"
	r.Params.TeamShort = teamShort
	r.Params.ChatShort = chatShort
	r.Params.ActorShort = actorShort
	r.Params.Uid = uid
	r.Params.Jid = chatShort.Jid
	r.Params.BuzzTimeout = int(timeout.Seconds())

	r.Params.Icons = chatShort.Icons

	// Set as Deprecated or just leave for backward compatibility?
	r.Params.Team = teamShort.Uid
	r.Params.DisplayName = chatShort.DisplayName
	return r
}

type ServerCallBuzz struct {
	BaseEvent
	Params serverCallBuzzParams `json:"params"`
}

type serverCallBuzzParams struct {
	TeamShort   TeamShort    `json:"teaminfo"`
	ChatShort   ChatShort    `json:"chat"`
	ActorShort  ContactShort `json:"actor"`
	Uid         string       `json:"uid"`
	Jid         JID          `json:"jid"`
	BuzzTimeout int          `json:"buzz_timeout"`
	Icons       *IconData    `json:"icons"`
	Team        string       `json:"team"`
	DisplayName string       `json:"display_name"`
}
