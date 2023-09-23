package command

type CreateUserCommand struct {
	Email          string
	PhotoUrl       string
	SocialId       string
	SocialUsername string
	Provider       string
}

type UpdateUserCommand struct {
	UserId   string
	Email    string
	Username string
	PhotoUrl string
}
