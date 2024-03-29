package tdproto

// Public Status
type PublicStatus struct {
	// Public Status Type
	Type PublicStatusType `json:"type"`

	// Display emoji
	Emoji string `json:"emoji"`

	// Status Label Russian
	StatusRu string `json:"status_ru"`

	// Status Label English
	StatusEn string `json:"status_en"`

	// Duration in seconds
	DurationSeconds int32 `json:"duration_seconds"`
	
	// Duration Label
	DurationLabel string `json:"duration_label"`
}

type ContactPublicStatus struct {
	// Public Status
	Status PublicStatus `json:"status"`

	// Expires at (iso datetime)
	ExpiresAt ISODateTimeString `json:"expires_at"`
}