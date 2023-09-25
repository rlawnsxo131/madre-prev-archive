package model

import (
	"database/sql"

	"time"

	"github.com/rlawnsxo131/madre-server/core/utils"
	"github.com/rlawnsxo131/madre-server/domain/entity/user"
)

type User struct {
	Id        int64          `db:"id"`
	Email     string         `db:"email"`
	Username  string         `db:"username"`
	PhotoUrl  sql.NullString `db:"photo_url"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}

type UserMapper struct{}

func (um UserMapper) MapToModel(u *user.User) *User {
	return &User{
		Id:        u.Id,
		Email:     u.Email,
		Username:  u.Username,
		PhotoUrl:  utils.NewNullString(u.PhotoUrl),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (um UserMapper) MapToEntity(u *User) *user.User {
	return &user.User{
		Id:        u.Id,
		Email:     u.Email,
		Username:  u.Username,
		PhotoUrl:  utils.NormalizeNullString(u.PhotoUrl),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
