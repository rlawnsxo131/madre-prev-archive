package user

type ReadRepository interface {
	FindOneById(id string) (*User, error)
	FindOneByUsername(username string) (*User, error)
}

type WriteRepository interface {
	Create(u *User) (string, error)
}

type ReadUseCase interface {
	ReadRepository
}

type WriteUseCase interface {
	WriteRepository
}
