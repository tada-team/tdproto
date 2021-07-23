package tdresp

import "github.com/tada-team/tdproto"

// Server response
type Resp struct {
	// Request status
	Ok bool `json:"ok"`

	// Result only if ok is true)
	Result interface{} `json:"result,omitempty"`

	// Error (only if ok is false)
	Error Err `json:"error,omitempty"`

	// Error (only if ok is false and Error is 'InvalidData')
	Details map[string]string `json:"details,omitempty"`

	// Reason (only if ok is false and Error is `AccessDenied`)
	Reason string `json:"reason,omitempty"`

	// Reason markup (only if ok is false and Error is `AccessDenied`)
	Markup []tdproto.MarkupEntity `json:"markup,omitempty"`

	// Server side work time
	DebugTime string `json:"_time,omitempty"`
}

type Err string

const (
	Ok                  = Err("OK")
	EmptyToken          = Err("EMPTY_TOKEN")
	EmptySession        = Err("EMPTY_SESSION")
	InvalidToken        = Err("INVALID_TOKEN")
	AccessDenied        = Err("ACCESS_DENIED")
	NotFound            = Err("NOT_FOUND")
	NotModified         = Err("NOT_MODIFIED")
	Gone                = Err("GONE")
	RateLimit           = Err("RATE_LIMIT")
	InternalServerError = Err("INTERNAL_SERVER_ERROR")
	InvalidMethod       = Err("INVALID_METHOD")
	InvalidData         = Err("INVALID_DATA")
)

func (e Err) Error() string { return string(e) }

func (e Err) StatusCode() int {
	switch e {
	case NotModified:
		return 304
	case EmptyToken, InvalidToken, EmptySession:
		return 401
	case AccessDenied:
		return 403
	case NotFound:
		return 404
	case Gone:
		return 410
	case RateLimit:
		return 429
	case InternalServerError:
		return 500
	case InvalidMethod:
		return 405
	case InvalidData:
		return 422
	default:
		return 200
	}
}
