package data

type writeUseCase struct {
	repo WriteRepository
}

func NewWriteUseCase() WriteUseCase {
	return &writeRepository{}
}
