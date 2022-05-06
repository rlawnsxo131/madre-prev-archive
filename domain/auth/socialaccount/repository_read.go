package socialaccount

import (
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

type ReadRepository interface {
	FindOneByProviderWithSocialId(provider string, socialId string) (*SocialAccount, error)
}

type readRepository struct {
	db     database.Database
	mapper entityMapper
}

func NewReadRepository(db database.Database) ReadRepository {
	return &readRepository{
		db:     db,
		mapper: entityMapper{},
	}
}

func (r *readRepository) FindOneByProviderWithSocialId(provider string, socialId string) (*SocialAccount, error) {
	var socialAccount SocialAccount

	query := "SELECT * FROM social_account" +
		" WHERE provider = $1 AND social_id = $2"

	err := r.db.QueryRowx(query, socialId, provider).StructScan(&socialAccount)
	if err != nil {
		customError := errors.Wrap(err, "ReadRepository: FindOneBySocialId")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return r.mapper.toEntity(&socialAccount), err
}
