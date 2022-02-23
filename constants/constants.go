package constants

type dbContextKey string

// http context
const (
	DBContextKey dbContextKey = "DB"
)

// http status
const (
	NotFoundErrorMessage       = "NOT_FOUND"             // 404
	InternalServerErrorMessage = "INTERNAL_SERVER_ERROR" // 500
)
