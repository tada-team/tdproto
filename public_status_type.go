package tdproto

// Public Status Types
type PublicStatusType string

const (
	PublicStatusNone PublicStatusType = "none"
	PublicStatusRemote PublicStatusType = "remote"
	PublicStatusVacation PublicStatusType = "vacation"
	PublicStatusSick PublicStatusType = "sick"
	PublicStatusCommuting PublicStatusType = "commuting"
	PublicStatusDoNotDisturb PublicStatusType = "do_not_disturb"
	PublicStatusLunch PublicStatusType = "lunch"
	PublicStatusBeRightBack PublicStatusType = "be_right_back"
)