package data

import "net/http"

type Controller interface {
	Get() http.HandlerFunc
	GetAll() http.HandlerFunc
}

type ReadUseCase interface {
	FindAll(limit int) ([]*Data, error)
	FindOneById(id string) (*Data, error)
}

type WriteUseCase interface {
	WriteRepository
}

type ReadRepository interface {
	FindAll(limit int) ([]*Data, error)
	FindOneById(id string) (*Data, error)
}

type WriteRepository interface{}
