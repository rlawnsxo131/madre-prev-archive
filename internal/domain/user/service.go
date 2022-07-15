package user

type UserQueryService interface {
	GetExistsSocialAccount(accessToken string) (bool, error)
}
