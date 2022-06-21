package datarepository

import "github.com/rlawnsxo131/madre-server-v3/internal/domain/data"

type dataQueryRepository struct{}

func NewDataQueryRepository() data.DataQueryRepository {
	return &dataQueryRepository{}
}
