package dataservice

import "github.com/rlawnsxo131/madre-server-v3/internal/domain/data"

type dataCommandService struct{}

func NewDataCommandService() data.DataCommandService {
	return &dataCommandService{}
}
