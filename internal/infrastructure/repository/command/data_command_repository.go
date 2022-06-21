package commandrepository

import "github.com/rlawnsxo131/madre-server-v3/internal/domain/data"

type dataCommandRepository struct{}

func NewDataCommandRepository() data.DataCommandRepository {
	return &dataCommandRepository{}
}
