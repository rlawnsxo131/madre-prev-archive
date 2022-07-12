package account

type QueryUsecase interface {
	CheckExistsSocialAccount(accessToken, provider string) (bool, error)
}

type CommandUsecase interface {
	CreateAccount(accessToken, username string) (*PublicAccount, error)
}
