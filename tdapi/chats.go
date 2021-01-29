package tdapi

type ChatFilter struct {
	UserParams
	Paginator

	// ?unread_first=true|false (default: false)
	UnreadFirst string `schema:"unread_first"`

	// ?unread_only=true|false (default: false)
	UnreadOnly string `schema:"unread_only"`

	// ?hidden=true|false|any (default: any)
	Hidden string `schema:"hidden"`

	// ?feed=true|false|any (default: any)
	Feed string `schema:"feed"`

	// ?has_messages=true|false|any (default: any)
	HasMessages string `schema:"has_messages"`

	// ?chat_type=task,group,direct|any (default: any)
	ChatType string `schema:"chat_type"`

	// ?mix_contacts=true|false (default: false)
	MixContacts string `schema:"mix_contacts"`

	// ?sort = [ "group_chat_display_name" | "-group_chat_display_name" | "last_message" | "-last_message" ]
	Sort string `schema:"sort"`

	// ?q=search-text
	Q string `schema:"q"`

	// gentime great than
	GentimeGT int64 `schema:"gentime_gt"`

	// deprecated
	HiddenOnly string `schema:"hidden_only"`
}
