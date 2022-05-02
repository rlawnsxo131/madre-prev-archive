package socialaccount

import (
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/database"
)

type WriteRepository interface {
	Create(socialAccount SocialAccount) (string, error)
}

type writeRepository struct {
	db database.Database
}

func NewWriteRepository(db database.Database) WriteRepository {
	return &writeRepository{
		db: db,
	}
}

func (r *writeRepository) Create(socialAccount SocialAccount) (string, error) {
	var id string
	var query = "INSERT INTO social_account(user_id, provider, social_id) VALUES(:user_id, :provider, :social_id) RETURNING id"

	err := r.db.PrepareNamedGet(&id, query, socialAccount)
	if err != nil {
		return "", errors.Wrap(err, "writeRepository: create")
	}

	return id, err
}
