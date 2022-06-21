package accountrepository

import (
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/external/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/account"
)

type accountCommandRepository struct {
	db     rdb.Database
	mapper accountMapper
}

func NewAccountCommandRepository(db rdb.Database) account.AccountCommandRepository {
	return &accountCommandRepository{db, accountMapper{}}
}

func (r *accountCommandRepository) InsertUser(u *account.User) (*account.User, error) {
	var id string

	query := "INSERT INTO public.user(email, origin_name, username, photo_url)" +
		" VALUES(:email, :origin_name, :username, :photo_url)" +
		" RETURNING id"

	err := r.db.PrepareNamedGet(
		&id,
		query,
		r.mapper.toUserModel(u),
	)
	if err != nil {
		return nil, errors.Wrap(err, "accountCommandRepository InsertUser")
	}
	u.ID = id

	return r.mapper.toUserEntity(u), nil
}

func (r *accountCommandRepository) InsertSocialAccount(sa *account.SocialAccount) (*account.SocialAccount, error) {
	var id string

	query := "INSERT INTO social_account(user_id, provider, social_id)" +
		" VALUES(:user_id, :provider, :social_id)" +
		" RETURNING id"

	err := r.db.PrepareNamedGet(
		&id,
		query,
		r.mapper.toSocialAccountModel(sa),
	)
	if err != nil {
		return nil, errors.Wrap(err, "accountCommandRepository InsertSocialAccount")
	}
	sa.ID = id

	return r.mapper.toSocialAccountEntity(sa), err
}
