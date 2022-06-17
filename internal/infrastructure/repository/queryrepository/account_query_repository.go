package queryrepository

import (
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/external/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/account"
	"github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/mapper"
	"github.com/rlawnsxo131/madre-server-v3/utils"
)

type accountQueryRepository struct {
	db     rdb.Database
	mapper mapper.AccountMapper
}

func NewAccountQueryRepository(db rdb.Database) account.AccountQueryRepository {
	return &accountQueryRepository{db, mapper.AccountMapper{}}
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

	return r.mapper.ToUserEntity(&u), err
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

	return r.mapper.ToUserEntity(&u), err
}

func (r *accountQueryRepository) ExistsUserByUsername(username string) (bool, error) {
	var exist bool

	query := "SELECT EXISTS" +
		"(SELECT 1 FROM public.user WHERE username = $1)"

	err := r.db.QueryRowx(query, username).Scan(&exist)
	if err != nil {
		customError := errors.Wrap(err, "accountQueryRepository ExistsUserByUsername")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return exist, err
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

	return r.mapper.ToSocialAccountEntity(&sa), err
}

func (r *accountQueryRepository) ExistsSocialAccountBySocialIdAndProvider(socialId, provider string) (bool, error) {
	var exist bool

	query := "SELECT EXISTS" +
		"(SELECT 1 FROM social_account WHERE social_id = $1 AND provider = $2)"

	err := r.db.QueryRowx(query, socialId, provider).Scan(&exist)
	if err != nil {
		customError := errors.Wrap(err, "accountQueryRepository ExistsSocialAccountBySocialIdAndProvider")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return exist, err
}
