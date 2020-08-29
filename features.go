package tdproto

// {hostname}/features.js / {hostname}/features.json structure
type Features struct {
	// current host
	Host string `json:"host"`

	// build/revision of server side
	Build string `json:"build"`

	// desktop application version
	DesktopVersion string `json:"desktop_version"`

	// webclient version
	FrontVersion string `json:"front_version"`

	// Application title
	AppTitle string `json:"app_title"`

	// Static files server address
	Userver string `json:"userver"`

	// Link to AppStore
	IOSApp string `json:"ios_app"`

	// Link to Google Play
	AndroidApp string `json:"android_app"`

	// Default UI theme
	Theme string `json:"theme"`

	// Minimal application version required for this server. Used for breaking changes
	MinAppVersion string `json:"min_app_version"`

	// Free registration allowed
	FreeRegistration bool `json:"free_registration"`

	// Maximum size of user's upload
	MaxUploadMb int `json:"max_upload_mb"`

	// Maximum number of forwarded messages
	MaxLinkedMessages int `json:"max_linked_messages"`

	// Maximum chars for text message
	MaxMessageLength int `json:"max_message_length"`

	// Maximum length for project and contact's sections names
	MaxSectionLength int `json:"max_section_length"`

	// Maximum length for tags
	MaxTagLength int `json:"max_tag_length"`

	// Maximum length for task title
	MaxTaskTitleLength int `json:"max_task_title_length"`

	// Maximum teams for one account
	MaxTeams int `json:"max_teams"`

	// Max inactivity seconds
	AfkAge int `json:"afk_age"`

	// Password authentication enabled
	AuthByPassword bool `json:"auth_by_password,omitempty"`

	// QR-code / link authentication enabled
	AuthByQrCode bool `json:"auth_by_qr_code,omitempty"`

	// SMS authentication enabled
	AuthBySms bool `json:"auth_by_sms,omitempty"`

	// Calls functions enabled
	Calls bool `json:"calls"`

	// Calls functions enabled for mobile applications
	MoblieCalls bool `json:"mobile_calls"`

	// Calls record enabled
	CallsRecord bool `json:"calls_record"`

	// ICE servers for WebRTC
	ICEServers []ICEServer `json:"ice_servers"`

	// True for onpremise installation
	CustomServer bool `json:"custom_server"`

	// name of instalation
	InstallationType string `json:"installation_type"`

	// testing installation
	IsTesting bool `json:"is_testing"`

	// yandex metrika counter id
	Metrika string `json:"metrika"`

	// minimal chars number for starting global search
	MinSearchLength int `json:"min_search_length"`

	// resend message in n seconds if no confirmation from server given
	ResendTimeout int `json:"resend_timeout"`

	// frontent sentry.io settings
	SentryDsnJS string `json:"sentry_dsn_js"`

	// message drafts saved on server
	ServerDrafts bool `json:"server_drafts"`

	// web-push notifacations related fields
	FirebaseAppId    string `json:"firebase_app_id"`
	FirebaseSenderId string `json:"firebase_sender_id"`
	SafariPushId     string `json:"safari_push_id"`

	// experimetal functions
	Terms                Terms `json:"terms"`
	SingleGroupTeams     bool  `json:"single_group_teams"`
	WikiPages            bool  `json:"wiki_pages"`
	AllowAdminMute       bool  `json:"allow_admin_mute,omitempty"`
	OnlyOneDevicePerCall bool  `json:"only_one_device_per_call,omitempty"`

	// obsolete fields. Always true
	TaskChecklist  bool `json:"task_checklist"`
	ReadonlyGroups bool `json:"readonly_groups"`
	TaskDashboard  bool `json:"task_dashboard"`
	TaskMessages   bool `json:"task_messages"`
	TaskPublic     bool `json:"task_public"`
	TaskTags       bool `json:"task_tags"`
}

type ICEServer struct {
	Urls string `json:"urls"`
}

type Terms struct {
	EnInTeam       string
	EnTeam         string
	EnTeamAccess   string
	EnTeamAdmin    string
	EnTeamAdmins   string
	EnTeamGuest    string
	EnTeamMember   string
	EnTeamMembers  string
	EnTeamOwner    string
	EnTeamSettings string
	RuTeamSettings string
	EnTeams        string
	EnToTeam       string
	RuInTeam       string
	RuTeam         string
	RuTeamAccess   string
	RuTeamAdmin    string
	RuTeamAdmins   string
	RuTeamD        string
	RuTeamGuest    string
	RuTeamMember   string
	RuTeamMembers  string
	RuTeamOwner    string
	RuTeamP        string
	RuTeamR        string
	RuTeams        string
	RuTeamsD       string
	RuTeamsP       string
	RuTeamsR       string
	RuTeamsT       string
	RuTeamsV       string
	RuTeamT        string
	RuTeamV        string
	RuToTeam       string
}
