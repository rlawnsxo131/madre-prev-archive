package user

type UserCommandRepository interface {
	Create(u *User) (string, error)
}

type UserQueryRepository interface {
	FindOneById(id string) (*User, error)
	FindOneByUsername(username string) (*User, error)
}
