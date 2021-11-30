package tdproto

// Personal account status
type PersonalAccountStatus string

const (
	// Personal account is active
	ActiveAccount PersonalAccountStatus = "PERSONAL_ACCOUNT_STATUS_ACTIVE"

	// Personal account is under financial blocking
	SuspendedAccount PersonalAccountStatus = "PERSONAL_ACCOUNT_STATUS_SUSPENDED"

	// Personal account is under administrative blocking
	BlockedAccount PersonalAccountStatus = "PERSONAL_ACCOUNT_STATUS_BLOCKED"
)
