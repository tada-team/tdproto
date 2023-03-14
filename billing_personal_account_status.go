package tdproto

// PersonalAccountStatus account status
type PersonalAccountStatus string

const (
	// ActiveAccount account is active
	ActiveAccount PersonalAccountStatus = "PERSONAL_ACCOUNT_STATUS_ACTIVE"

	// SuspendedAccount account is under financial blocking
	SuspendedAccount PersonalAccountStatus = "PERSONAL_ACCOUNT_STATUS_SUSPENDED"

	// BlockedAccount account is under administrative blocking
	BlockedAccount PersonalAccountStatus = "PERSONAL_ACCOUNT_STATUS_BLOCKED"

	// UnspecifiedStatus Unknown account status
	UnspecifiedStatus PersonalAccountStatus = "PERSONAL_ACCOUNT_STATUS_UNSPECIFIED"
)
