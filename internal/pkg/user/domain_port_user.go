package user

type UserService interface {
	Create(u *User) (string, error)
	FindOneById(id string) (*User, error)
	FindOneByUsername(username string) (*User, error)
}

type UserRepository interface {
	Create(u *User) (string, error)
	FindOneById(id string) (*User, error)
	FindOneByUsername(username string) (*User, error)
}
