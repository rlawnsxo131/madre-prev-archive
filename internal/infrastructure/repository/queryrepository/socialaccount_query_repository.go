package queryrepository

import (
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/external/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/account"
	"github.com/rlawnsxo131/madre-server-v3/utils"
)

type socialAccountQueryRepository struct {
	db rdb.Database
}

func NewSocialAccountQueryRepository(db rdb.Database) account.SocialAccountQueryRepository {
	return &socialAccountQueryRepository{db}
}

func (r *socialAccountQueryRepository) FindOneBySocialIdAndProvider(socialId, provider string) (*account.SocialAccount, error) {
	var sa account.SocialAccount

	query := "SELECT * FROM social_account" +
		" WHERE social_id = $1" +
		" AND provider = $2"

	err := r.db.QueryRowx(query, socialId, provider).StructScan(&sa)
	if err != nil {
		customError := errors.Wrap(err, "socialaccount ReadRepository FindOneBySocialId")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return &sa, err
}
