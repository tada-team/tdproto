package tdproto

type PushDevice struct {
	Type                  string `json:"type"`
	DeviceId              string `json:"device_id"`
	NotificationToken     string `json:"notification_token"`
	VoipNotificationToken string `json:"voip_notification_token"`
	AllowedNotifications  bool   `json:"allowed_notifications"` // deprecated
	Name                  string `json:"name"`
	DataPushes            bool   `json:"data_pushes"`
	DataBadges            bool   `json:"data_badges"`
}
