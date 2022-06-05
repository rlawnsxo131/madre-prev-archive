package data

type ReadRepository interface {
	FindAll(limit int) ([]*Data, error)
	FindOneById(id string) (*Data, error)
}

type WriteRepository interface{}

type ReadUseCase interface {
	ReadRepository
}

type WriteUseCase interface {
	WriteRepository
}
