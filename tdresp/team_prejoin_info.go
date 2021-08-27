package tdresp

import "github.com/tada-team/tdproto"

// Invitation information
type TeamPrejoinInfo struct {
	// Short team information
	Team tdproto.TeamShort `json:"team"`
}
