package tdproto

// Personal account status
type PersonalAccountStatus string

const (
	// Personal account is active
	ActiveAccount PersonalAccountStatus = "Active"

	// Personal account is under financial blocking
	SuspendedAccount PersonalAccountStatus = "Suspended"

	// Personal account is under administrative blocking
	BlockedAccount PersonalAccountStatus = "Blocked"
)
