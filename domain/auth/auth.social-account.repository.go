package auth

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type SocialAccountRepository interface {
	FindOneBySocialId(socialId string) (SocialAccount, error)
}

type socialAccountRepository struct {
	db *sqlx.DB
}

func NewSocialAccountRepository(db *sqlx.DB) *socialAccountRepository {
	return &socialAccountRepository{
		db: db,
	}
}

func (r *socialAccountRepository) FindOneBySocialId(socialId string) (SocialAccount, error) {
	var socialAccount SocialAccount

	sql := "SELECT * FROM social_account WHERE social_id = ?"
	err := r.db.QueryRowx(sql, socialId).StructScan(&socialAccount)
	if err != nil {
		err = errors.Wrap(err, "SocialAccountRepository: FindOneBySocialId sql error")
	}

	return socialAccount, err
}
