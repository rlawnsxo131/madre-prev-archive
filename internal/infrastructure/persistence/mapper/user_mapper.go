package mapper

import (
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
	"github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/persistence/model"
	"github.com/rlawnsxo131/madre-server-v3/lib/utils"
)

type UserMapper struct{}

func (um UserMapper) MapToModel(u *user.User) *model.User {
	return &model.User{
		Id:        u.Id,
		Email:     u.Email,
		Username:  u.Email,
		PhotoUrl:  utils.NewNullString(u.PhotoUrl),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (um UserMapper) MapToEntity(u *model.User) *user.User {
	return &user.User{
		Id:        u.Id,
		Email:     u.Email,
		Username:  u.Email,
		PhotoUrl:  utils.NormalizeNullString(u.PhotoUrl),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
