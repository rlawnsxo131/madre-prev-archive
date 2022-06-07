package socialaccount

import (
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

type ReadRepository interface {
	FindOneBySocialIdAndProvider(params *SocialIDAndProviderDto) (*SocialAccount, error)
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

func (r *readRepository) FindOneBySocialIdAndProvider(params *SocialIDAndProviderDto) (*SocialAccount, error) {
	var sa SocialAccount

	query := "SELECT * FROM social_account" +
		" WHERE social_id = :social_id" +
		" AND provider = :provider"

	err := r.db.PrepareNamedGet(
		&sa,
		query,
		params,
	)
	if err != nil {
		customError := errors.Wrap(err, "socialaccount ReadRepository FindOneBySocialId")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return r.mapper.toEntity(&sa), err
}
