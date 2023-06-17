package mapper

import (
	valueutil "github.com/rlawnsxo131/madre-server-v3/core/utils/value-util"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
	"github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/persistence/model"
)

type UserMapper struct{}

func (um UserMapper) MapToModel(u *user.User) *model.User {
	return &model.User{
		Id:        u.Id,
		Email:     u.Email,
		Username:  u.Username,
		PhotoUrl:  valueutil.NewNullString(u.PhotoUrl),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (um UserMapper) MapToEntity(u *model.User) *user.User {
	return &user.User{
		Id:        u.Id,
		Email:     u.Email,
		Username:  u.Username,
		PhotoUrl:  valueutil.NormalizeNullString(u.PhotoUrl),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
