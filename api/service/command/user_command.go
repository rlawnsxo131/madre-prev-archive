package command

type CreateUserCommand struct {
	Provider    string
	AccessToken string
}

type UpdateUsernameCommand struct {
	Username string
}
