package dataservice

import "github.com/rlawnsxo131/madre-server-v3/internal/domain/data"

type dataQueryService struct{}

func NewDataQueryService() data.DataQueryService {
	return &dataQueryService{}
}
