package data

type WriteUseCase interface {
	WriteRepository
}

type writeUseCase struct {
	repo WriteRepository
}

func NewWriteUseCase() WriteUseCase {
	return &writeRepository{}
}
