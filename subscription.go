package tdproto

import "time"

// Subscription - an entity that signifies the fact
// of subscribing to the tariff of any team for a certain period
// (not defined, in the case of the default tariff)
type Subscription struct {
	// Subscription id
	Uid string `json:"uid"`
	// Subscription activation time
	Activated *time.Time `json:"activated,omitempty"`
	// Subscription expiration time
	Expires *time.Time `json:"expires,omitempty"`
	// ID of the tariff for which the subscription is valid
	TariffUid string `json:"tariff_uid,omitempty"`
	// ID of the user who subscribed
	UserUid string `json:"user_uid,omitempty"`
	// EmptyWorkplaceCount empty workplace count
	EmptyWorkplaceCount uint32 `json:"empty_workplace_count,omitempty"`
}
