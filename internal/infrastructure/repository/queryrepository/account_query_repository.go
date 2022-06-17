package queryrepository

import (
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/external/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/account"
	"github.com/rlawnsxo131/madre-server-v3/utils"
)

type accountQueryRepository struct {
	db rdb.Database
}

func NewAccountQueryRepository(db rdb.Database) account.AccountQueryRepository {
	return &accountQueryRepository{db}
}

func (r *accountQueryRepository) FindUserById(id string) (*account.User, error) {
	var u account.User

	query := "SELECT * FROM public.user" +
		" WHERE id = $1"

	err := r.db.QueryRowx(query, id).StructScan(&u)
	if err != nil {
		customError := errors.Wrap(err, "accountQueryRepository FindUserById")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return &u, err
}

func (r *accountQueryRepository) FindUserByUsername(username string) (*account.User, error) {
	var u account.User

	query := "SELECT * FROM public.user" +
		" WHERE username = $1"

	err := r.db.QueryRowx(query, username).StructScan(&u)
	if err != nil {
		customError := errors.Wrap(err, "accountQueryRepository FindUserUsername")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return &u, err
}

func (r *accountQueryRepository) FindSocialAccountBySocialIdAndProvider(socialId, provider string) (*account.SocialAccount, error) {
	var sa account.SocialAccount

	query := "SELECT * FROM social_account" +
		" WHERE social_id = $1" +
		" AND provider = $2"

	err := r.db.QueryRowx(query, socialId, provider).StructScan(&sa)
	if err != nil {
		customError := errors.Wrap(err, "accountQueryRepository FindSocialAccountBySocialIdAndProvider")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return &sa, err
}
