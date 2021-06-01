package openapi

type Format string

const (
	Uuid     Format = "uuid"
	Binary   Format = "binary"
	Email    Format = "email"
	Int32    Format = "int32"
	Int64    Format = "int64"
	Base64   Format = "base64"
	DateTime Format = "date-time"
)
