package constants

type dbContextKey string

// http context
const (
	DBContextKey dbContextKey = "DB"
)

// http status
const (
	ErrNotFoundMessage       = "NOT_FOUND"             // 404
	ErrInternalServerMessage = "INTERNAL_SERVER_ERROR" // 500
)
