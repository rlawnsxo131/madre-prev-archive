package auth

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type SocialAccountReadRepository interface {
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

func (r *socialAccountReadRepository) FindOneBySocialId(socialId string) (SocialAccount, error) {
	var socialAccount SocialAccount

	sql := "SELECT * FROM social_account WHERE social_id = ?"
	err := r.db.QueryRowx(sql, socialId).StructScan(&socialAccount)
	if err != nil {
		err = errors.Wrap(err, "SocialAccountRepository: FindOneBySocialId sql error")
	}

	return socialAccount, err
}
