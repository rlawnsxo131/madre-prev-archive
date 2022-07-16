package mapper

import (
	"database/sql"
	"time"

	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
	"github.com/rlawnsxo131/madre-server-v3/utils"
)

type UserModel struct {
	Id        string         `db:"id"`
	Email     string         `db:"email"`
	Username  string         `db:"username"`
	PhotoUrl  sql.NullString `db:"photo_url"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}

type UserMapper struct{}

func (um UserMapper) MapToModel(u *user.User) *UserModel {
	return &UserModel{
		Id:        u.Id,
		Email:     u.Email,
		Username:  u.Email,
		PhotoUrl:  utils.NewNullString(u.PhotoUrl),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (um UserMapper) MapToEntity(u *UserModel) *user.User {
	return &user.User{
		Id:        u.Id,
		Email:     u.Email,
		Username:  u.Email,
		PhotoUrl:  utils.NormalizeNullString(u.PhotoUrl),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
