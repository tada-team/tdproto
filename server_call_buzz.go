package tdproto

import (
	"time"
)

func ServerCallBuzz(teamShort TeamShort, chatShort ChatShort, actorShort ContactShort, uid string, timeout time.Duration) (resp struct {
	BaseEvent
	Params struct {
		TeamShort   TeamShort    `json:"teaminfo"`
		ChatShort   ChatShort    `json:"chat"`
		ActorShort  ContactShort `json:"actor"`
		Uid         string       `json:"uid"`
		Jid         JID          `json:"jid"`
		BuzzTimeout int          `json:"buzz_timeout"`
		Icons       *IconData    `json:"icons"`

		Team        string `json:"team"`
		DisplayName string `json:"display_name"`
	} `json:"params"`
}) {

	resp.BaseEvent.Name = "server.call.buzz"
	resp.Params.TeamShort = teamShort
	resp.Params.ChatShort = chatShort
	resp.Params.ActorShort = actorShort
	resp.Params.Uid = uid
	resp.Params.Jid = chatShort.Jid
	resp.Params.BuzzTimeout = int(timeout.Seconds())

	resp.Params.Icons = chatShort.Icons

	// Set as Deprecated or just leave for reverse compatibility?
	resp.Params.Team = teamShort.Uid
	resp.Params.DisplayName = chatShort.DisplayName
	return resp
}