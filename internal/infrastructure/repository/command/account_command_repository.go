package commandrepository

import (
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/external/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/account"
	"github.com/rlawnsxo131/madre-server-v3/internal/infrastructure"
)

type accountCommandRepository struct {
	db     rdb.Database
	mapper infrastructure.AccountMapper
}

func NewAccountCommandRepository(db rdb.Database) account.AccountCommandRepository {
	return &accountCommandRepository{db, infrastructure.AccountMapper{}}
}

func (r *accountCommandRepository) InsertUser(u *account.User) (string, error) {
	var id string

	query := "INSERT INTO public.user(email, origin_name, username, photo_url)" +
		" VALUES(:email, :origin_name, :username, :photo_url)" +
		" RETURNING id"

	err := r.db.PrepareNamedGet(
		&id,
		query,
		r.mapper.ToUserModel(u),
	)
	if err != nil {
		return "", errors.Wrap(err, "accountCommandRepository InsertUser")
	}

	return id, nil
}

func (r *accountCommandRepository) InsertSocialAccount(sa *account.SocialAccount) (string, error) {
	var id string

	query := "INSERT INTO social_account(user_id, provider, social_id)" +
		" VALUES(:user_id, :provider, :social_id)" +
		" RETURNING id"

	err := r.db.PrepareNamedGet(
		&id,
		query,
		r.mapper.ToSocialAccountModel(sa),
	)
	if err != nil {
		return "", errors.Wrap(err, "accountCommandRepository InsertSocialAccount")
	}

	return id, err
}
