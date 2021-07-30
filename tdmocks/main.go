package main

import (
	"github.com/tada-team/tdproto"
)

var Date = "2019-09-18T00:00:07.435409Z"

var Gentime int64 = 1626960084072284400

var Section = tdproto.Section{
	Uid:          "00000000-0000-0000-0000-000000000001",
	SortOrdering: 0,
	Name:         "Managers",
	Gentime:      Gentime,
	Description:  "All managers",
}

var StubIcons = tdproto.IconData{
	Letters: "AS",
	Color:   "#e3665a",
	Sm: tdproto.SingleIcon{
		Url:    "https://web.tada.team/a/e3665a:as/256.png",
		Width:  256,
		Height: 256,
	},
	Lg: tdproto.SingleIcon{
		Url:    "https://web.tada.team/a/e3665a:as/512.png",
		Width:  512,
		Height: 512,
	},
}

var Alice = tdproto.Contact{
	Jid:              "d-00000000-0000-0000-0000-000000000001",
	DisplayName:      "Alice Smith",
	ShortName:        "Alice S.",
	ContactEmail:     "alice@example.com",
	ContactPhone:     "+7 (555) 000-00-01",
	Icons:            StubIcons,
	Gentime:          Gentime,
	Role:             "manager",
	Mood:             "🤦",
	TeamStatus:       tdproto.TeamMember,
	LastActivity:     Date,
	Sections:         []string{Section.Uid},
	CanSendMessage:   true,
	CanCall:          true,
	CanCreateTask:    true,
	CanImportTasks:   true,
	CanAddToGroup:    true,
	CanDelete:        true,
	ChangeableFields: []string{"role"},
	FamilyName:       "Smith",
	GivenName:        "Alice",
}
