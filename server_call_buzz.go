package tdproto

import (
	"time"
)

func NewServerCallBuzz(chat JID, teamUid, displayName string, icons *IconData, uid string, timeout time.Duration) (r ServerCallBuzz) {
	r.BaseEvent.Name = "server.call.buzz"
	r.Params.Jid = chat
	r.Params.Team = teamUid
	r.Params.DisplayName = displayName
	r.Params.Icons = icons
	r.Params.Uid = uid
	r.Params.BuzzTimeout = int(timeout.Seconds())
	return r
}

type ServerCallBuzz struct {
	BaseEvent
	Params serverCallBuzzParams `json:"params"`
}

type serverCallBuzzParams struct {
	Jid         JID       `json:"jid"`
	Team        string    `json:"team"`
	DisplayName string    `json:"display_name"`
	Icons       *IconData `json:"icons"`
	Uid         string    `json:"uid"`
	BuzzTimeout int       `json:"buzz_timeout"`
}
