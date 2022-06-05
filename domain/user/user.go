package user

import "net/http"

type Controller interface {
	Get() http.HandlerFunc
	Put() http.HandlerFunc
}

type ReadUseCase interface {
	FindOneById(id string) (*User, error)
	FindOneByUsername(username string) (*User, error)
}

type WriteUseCase interface {
	Create(u *User) (string, error)
}

type ReadRepository interface {
	FindOneById(id string) (*User, error)
	FindOneByUsername(username string) (*User, error)
}

type WriteRepository interface {
	Create(u *User) (string, error)
}
