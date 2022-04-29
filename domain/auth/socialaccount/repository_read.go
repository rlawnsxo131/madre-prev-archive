package socialaccount

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

type SocialAccountReadRepository interface {
	FindOneByProviderWithSocialId(provider string, socialId string) (*SocialAccount, error)
}

type socialAccountReadRepository struct {
	ql logger.QueryLogger
}

func NewSocialAccountReadRepository(db *sqlx.DB) SocialAccountReadRepository {
	return &socialAccountReadRepository{
		ql: logger.NewQueryLogger(db),
	}
}

func (r *socialAccountReadRepository) FindOneByProviderWithSocialId(provider string, socialId string) (*SocialAccount, error) {
	var socialAccount SocialAccount

	query := "SELECT * FROM social_account WHERE provider = $1 AND social_id = $2"
	err := r.ql.QueryRowx(query, socialId, provider).StructScan(&socialAccount)
	if err != nil {
		customError := errors.Wrap(err, "SocialAccountReadRepository: FindOneBySocialId")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return &socialAccount, err
}
