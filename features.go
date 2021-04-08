package tdproto

const (
	SupportedFeaturesHeader = "X-Supported-Features"
	Feature2FactorAuth      = "2FactorAuth"
)

// Server information. Readonly.
type Features struct {
	// Current host
	Host string `json:"host"`

	// Build/revision of server side
	Build string `json:"build"`

	// Desktop application version
	DesktopVersion string `json:"desktop_version"`

	// Webclient version
	FrontVersion string `json:"front_version"`

	// Application title
	AppTitle string `json:"app_title"`

	// Landing page address, if any
	LandingUrl string `json:"landing_url,omitempty"`

	// Local applications urls
	AppSchemes []string `json:"app_schemes"`

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

	// Maximum number of message uploads
	MaxMessageUploads int `json:"max_message_uploads"`

	// Maximum chars for: family_name, given_name, patronymic if any
	MaxUsernamePartLength int `json:"max_username_part_length"`

	// Maximum chars for group chat name
	MaxGroupTitleLength int `json:"max_group_title_length"`

	// Maximum chars for role in team
	MaxRoleLength int `json:"max_role_length"`

	// Maximum chars for mood in team
	MaxMoodLength int `json:"max_mood_length"`

	// Maximum chars for text message
	MaxMessageLength int `json:"max_message_length"`

	// Maximum length for project and contact's sections names
	MaxSectionLength int `json:"max_section_length"`

	// Maximum length for tags
	MaxTagLength int `json:"max_tag_length"`

	// Maximum length for task title
	MaxTaskTitleLength int `json:"max_task_title_length"`

	// Maximum length for ColorRule description
	MaxColorRuleDescriptionLength int `json:"max_color_rule_description_length"`

	//Maximum length for urls
	MaxUrlLength int `json:"max_url_length"`

	// Maximum length for Integration comment
	MaxIntegrationCommentLength int `json:"max_integration_comment_length"`

	// Maximum teams for one account
	MaxTeams int `json:"max_teams"`

	// Maximum search result
	MaxMessageSearchLimit int `json:"max_message_search_limit"`

	// Max inactivity seconds
	AfkAge int `json:"afk_age"`

	// Password authentication enabled
	AuthByPassword bool `json:"auth_by_password,omitempty"`

	// QR-code / link authentication enabled
	AuthByQrCode bool `json:"auth_by_qr_code,omitempty"`

	// SMS authentication enabled
	AuthBySms bool `json:"auth_by_sms,omitempty"`

	// Two-factor authentication (2FA) enabled
	Auth2fa bool `json:"auth_2fa,omitempty"`

	// External services
	OAuthServices []OAuthService `json:"oauth_services,omitempty"`

	// ICE servers for WebRTC
	ICEServers []ICEServer `json:"ice_servers"`

	// True for premise installation
	CustomServer bool `json:"custom_server"`

	// Name of installation
	InstallationType string `json:"installation_type"`

	// Testing installation
	IsTesting bool `json:"is_testing"`

	// Yandex metrika counter id
	Metrika string `json:"metrika"`

	// Minimal chars number for starting global search
	MinSearchLength int `json:"min_search_length"`

	// Resend message in n seconds if no confirmation from server given
	ResendTimeout int `json:"resend_timeout"`

	// Frontend sentry.io settings
	SentryDsnJS string `json:"sentry_dsn_js"`

	// Message drafts saved on server
	ServerDrafts bool `json:"server_drafts"`

	// Firebase settings for web-push notifications
	FirebaseAppId string `json:"firebase_app_id"`

	// Firebase settings for web-push notifications
	FirebaseSenderId string `json:"firebase_sender_id"`

	// Firebase settings for web-push notifications
	FirebaseApiKey string `json:"firebase_api_key"`

	// Firebase settings for web-push notifications
	FirebaseAuthDomain string `json:"firebase_auth_domain"`

	// Firebase settings for web-push notifications
	FirebaseDatabaseUrl string `json:"firebase_database_url"`

	// Firebase settings for web-push notifications
	FirebaseProjectId string `json:"firebase_project_id"`

	// Firebase settings for web-push notifications
	FirebaseStorageBucket string `json:"firebase_storage_bucket"`

	// Calls functions enabled
	Calls bool `json:"calls"`

	// Calls functions enabled for mobile applications
	MobileCalls bool `json:"mobile_calls"`

	// Calls record enabled
	CallsRecord bool `json:"calls_record"`

	// Disallow call from multiply devices. Experimental
	OnlyOneDevicePerCall bool `json:"only_one_device_per_call,omitempty"`

	// Maximum number of participants per call
	MaxParticipantsPerCall int `json:"max_participants_per_call,omitempty"`

	// Safari push id for web-push notifications
	SafariPushId string `json:"safari_push_id"`

	// Multiple message uploads
	MessageUploads bool `json:"message_uploads"`

	// Team entity naming. Experimental.
	Terms Terms `json:"terms"`

	// Cross team communication. Experimental.
	SingleGroupTeams bool `json:"single_group_teams"`

	// Wiki pages in chats. Experimental
	WikiPages bool `json:"wiki_pages"`

	// Wiki pages in chats. Experimental
	AllowAdminMute bool `json:"allow_admin_mute,omitempty"`

	// Default wallpaper url for mobile apps, if any
	DefaultWallpaper *Wallpaper `json:"default_wallpaper,omitempty"`

	// Deprecated
	TaskChecklist bool `json:"task_checklist"`

	// Deprecated
	ReadonlyGroups bool `json:"readonly_groups"`

	// Deprecated
	TaskDashboard bool `json:"task_dashboard"`

	// Deprecated
	TaskMessages bool `json:"task_messages"`

	// Deprecated
	TaskPublic bool `json:"task_public"`

	// Deprecated
	TaskTags bool `json:"task_tags"`
}

// Interactive Connectivity Establishment Server for WEB Rtc connection. Readonly.
type ICEServer struct {
	// URls
	Urls string `json:"urls"`
}

// Experimental translation fields for "team" entity renaming. Readonly.
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
