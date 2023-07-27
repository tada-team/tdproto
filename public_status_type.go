package tdproto

// Public Status Types
type PublicStatusType string

const (
	None PublicStatusType = "none"
	Remote PublicStatusType = "remote"
	Vacation PublicStatusType = "vacation"
	Sick PublicStatusType = "sick"
	Commuting PublicStatusType = "commuting"
	DoNotDisturb PublicStatusType = "do_not_disturb"
	Lunch PublicStatusType = "lunch"
	BeRightBack PublicStatusType = "be_right_back"
)