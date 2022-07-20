package tdproto

// ActionType for server.processing event
type ActionType string

const (
	// ActionTypeContactImport ...
	ActionTypeContactImport ActionType = "contact_import"

	// ActionTypeTaskImport ...
	ActionTypeTaskImport ActionType = "task_import"

	// ActionTypeArchiveUnpacking ...
	ActionTypeArchiveUnpacking ActionType = "archive_unpacking"

	// ActionTypeGenerateChat ...
	ActionTypeGenerateChat ActionType = "generate_chat"
)

func (at ActionType) String() string { return string(at) }
