package auth

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

type AuthTokenReadRepository interface {
	FindOneByUserId(userId int64) (AuthToken, error)
}

type authTokenReadRepository struct {
	db *sqlx.DB
}

func NewAuthTokenReadRepository(db *sqlx.DB) AuthTokenReadRepository {
	return &authTokenReadRepository{
		db: db,
	}
}

func (r *authTokenReadRepository) FindOneByUserId(userId int64) (AuthToken, error) {
	var authToken AuthToken

	query := "SELECT * FROM auth_token WHERE user_id = ? "

	err := r.db.QueryRowx(query, userId).StructScan(&authToken)
	if err != nil {
		customError := errors.Wrap(err, "SocialAccountRepository: FindOneById")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return authToken, nil
}
