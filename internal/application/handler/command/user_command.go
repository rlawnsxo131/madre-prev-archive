package command

type CreateUserCommand struct {
	Email          string
	Username       string
	PhotoUrl       string
	SocialId       string
	SocialUsername string
	Provider       string
}

type UpdateUserCommand struct {
	Email    string
	Username string
	PhotoUrl string
}
