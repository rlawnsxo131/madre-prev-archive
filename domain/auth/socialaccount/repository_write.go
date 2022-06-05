package socialaccount

import (
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/database"
)

type writeRepository struct {
	db     database.Database
	mapper entityMapper
}

func NewWriteRepository(db database.Database) WriteRepository {
	return &writeRepository{
		db:     db,
		mapper: entityMapper{},
	}
}

func (r *writeRepository) Create(sa *SocialAccount) (string, error) {
	var id string

	query := "INSERT INTO social_account(user_id, provider, social_id)" +
		" VALUES(:user_id, :provider, :social_id)" +
		" RETURNING id"

	err := r.db.PrepareNamedGet(
		&id,
		query,
		r.mapper.toModel(sa),
	)
	if err != nil {
		return "", errors.Wrap(err, "socialaccount WriteRepository create")
	}

	return id, err
}
