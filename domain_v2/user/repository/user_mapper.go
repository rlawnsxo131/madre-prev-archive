package repository

import (
	"github.com/rlawnsxo131/madre-server-v2/domain_v2/user"
)

type userEntityMapper struct{}

func (e userEntityMapper) toEntity(u *user.User) *user.User {
	return &user.User{
		ID:         u.ID,
		Email:      u.Email,
		OriginName: u.OriginName,
		Username:   u.Username,
		PhotoUrl:   u.PhotoUrl,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

func (e userEntityMapper) toModel(u *user.User) *user.User {
	return &user.User{
		Email:      u.Email,
		OriginName: u.OriginName,
		Username:   u.Username,
		PhotoUrl:   u.PhotoUrl,
		UpdatedAt:  u.UpdatedAt,
	}
}
