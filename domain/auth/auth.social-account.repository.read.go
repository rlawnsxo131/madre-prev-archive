package auth

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

type SocialAccountReadRepository interface {
	FindOneById(id int64) (SocialAccount, error)
	FindOneBySocialIdWithProvider(socialId string, provider string) (SocialAccount, error)
}

type socialAccountReadRepository struct {
	db *sqlx.DB
}

func NewSocialAccountReadRepository(db *sqlx.DB) SocialAccountReadRepository {
	return &socialAccountReadRepository{
		db: db,
	}
}

func (r *socialAccountReadRepository) FindOneById(id int64) (SocialAccount, error) {
	var socialAccount SocialAccount

	query := "SELECT * FROM social_account WHERE id = ?"
	err := r.db.QueryRowx(query, id).StructScan(&socialAccount)
	if err != nil {
		customError := errors.Wrap(err, "SocialAccountRepository: FindOneById")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return socialAccount, err
}

func (r *socialAccountReadRepository) FindOneBySocialIdWithProvider(socialId string, provider string) (SocialAccount, error) {
	var socialAccount SocialAccount

	query := "SELECT * FROM social_account WHERE social_id = ? AND provider = ?"
	err := r.db.QueryRowx(query, socialId, provider).StructScan(&socialAccount)
	if err != nil {
		customError := errors.Wrap(err, "SocialAccountRepository: FindOneBySocialId")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return socialAccount, err
}
