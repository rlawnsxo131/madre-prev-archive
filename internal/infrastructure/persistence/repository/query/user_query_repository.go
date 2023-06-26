package queryrepository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rlawnsxo131/madre-server-v3/core/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
	"github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/persistence/mapper"
	"github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/persistence/model"
)

type userQueryRepository struct {
	db     rdb.SingletonDatabase
	mapper mapper.UserMapper
}

func NewUserQueryRepository(db rdb.SingletonDatabase) user.UserQueryRepository {
	return &userQueryRepository{
		db:     db,
		mapper: mapper.UserMapper{},
	}
}

func (uqr *userQueryRepository) FindById(id string) (*user.User, error) {
	conn, err := uqr.db.Conn()
	if err != nil {
		return nil, err
	}

	u := model.User{}
	query := "SELECT id, email, username, photo_url, created_at, updated_at" +
		" FROM public.user" +
		" WHERE id = @id"

	row := conn.QueryRow(context.Background(), query, pgx.NamedArgs{
		"id": id,
	})
	err = row.Scan(
		&u.Id,
		&u.Email,
		&u.Username,
		&u.PhotoUrl,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return uqr.mapper.MapToEntity(&u), nil
}

func (uqr *userQueryRepository) FindByUsername(username string) (*user.User, error) {
	u := model.User{}

	return uqr.mapper.MapToEntity(&u), nil
}

func (uqr *userQueryRepository) ExistsByUsername(username string) (bool, error) {
	return false, nil
}
