package auth

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/lib"
)

var sqlxManager = lib.GetSqlxManager()

type SocialAccountReadRepository interface {
	FindOneById(id int64) (SocialAccount, error)
	FindOneBySocialId(socialId string) (SocialAccount, error)
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
		customError := errors.Wrap(err, "SocialAccountRepository: FindOneById error")
		err = sqlxManager.ErrNoRowsReturnRawError(err, customError)
	}

	return socialAccount, err
}

func (r *socialAccountReadRepository) FindOneBySocialId(socialId string) (SocialAccount, error) {
	var socialAccount SocialAccount

	query := "SELECT * FROM social_account WHERE social_id = ?"
	err := r.db.QueryRowx(query, socialId).StructScan(&socialAccount)
	if err != nil {
		customError := errors.Wrap(err, "SocialAccountRepository: FindOneBySocialId error")
		err = sqlxManager.ErrNoRowsReturnRawError(err, customError)
	}

	return socialAccount, err
}
