package tdapi

type Err string

const (
	Ok                  = Err("OK")
	EmptyToken          = Err("EMPTY_TOKEN")
	EmptySession        = Err("EMPTY_SESSION")
	InvalidToken        = Err("INVALID_TOKEN")
	AccessDenied        = Err("ACCESS_DENIED")
	NotFound            = Err("NOT_FOUND")
	NotModified         = Err("NOT_MODIFIED")
	RateLimit           = Err("RATE_LIMIT")
	InternalServerError = Err("INTERNAL_SERVER_ERROR")
	InvalidMethod       = Err("INVALID_METHOD")
	InvalidData         = Err("INVALID_DATA")
	AccountNotFound     = Err("ACCOUNT_NOT_FOUND")
	AccountSuspended    = Err("ACCOUNT_SUSPENDED")
	AccountBlocked      = Err("ACCOUNT_BLOCKED")
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
	case RateLimit:
		return 429
	case InternalServerError:
		return 500
	case InvalidMethod:
		return 405
	case InvalidData:
		return 422
	case AccountNotFound, AccountBlocked:
		return 402
	case AccountSuspended:
		return 451
	default:
		return 200
	}
}
