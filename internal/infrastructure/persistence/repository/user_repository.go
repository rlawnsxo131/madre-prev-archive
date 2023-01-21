package repository

import (
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
)

type userRepository struct{}

func NewUserRepository() user.UserRepository {
	return &userRepository{}
}

const (
	INSERT_USER_SQL = "INSERT INTO public.user(email, origin_name, username, photo_url)" +
		" VALUES(:email, :username, :photo_url)" +
		" RETURNING id"

	INSERT_SOCIAL_ACCOUNT_SQL = "INSERT INTO public.social_account(user_id, social_id, social_username, provider)" +
		" VALUES(:user_id, :social_id, :social_username, :provider)" +
		" RETURNING id"
)

func (ur *userRepository) Create(u *user.User) (string, error)
