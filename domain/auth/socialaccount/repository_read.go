package socialaccount

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

type ReadRepository interface {
	FindOneByProviderWithSocialId(provider string, socialId string) (*SocialAccount, error)
}

type readRepository struct {
	ql logger.QueryLogger
}

func NewReadRepository(db *sqlx.DB) ReadRepository {
	return &readRepository{
		ql: logger.NewQueryLogger(db),
	}
}

func (r *readRepository) FindOneByProviderWithSocialId(provider string, socialId string) (*SocialAccount, error) {
	var socialAccount SocialAccount

	query := "SELECT * FROM social_account WHERE provider = $1 AND social_id = $2"
	err := r.ql.QueryRowx(query, socialId, provider).StructScan(&socialAccount)
	if err != nil {
		customError := errors.Wrap(err, "ReadRepository: FindOneBySocialId")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return &socialAccount, err
}
