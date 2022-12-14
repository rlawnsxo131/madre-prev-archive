package query

type UserQueryHandler interface{}

type userQueryHandler struct{}

func NewUserQueryHandler() UserQueryHandler {
	return &userQueryHandler{}
}
