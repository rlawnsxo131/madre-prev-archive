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
	sql := "SELECT id, username" +
		" FROM public.user" +
		" WHERE id = @id" +
		" AND username = @username"

	row := conn.QueryRow(
		context.TODO(),
		sql,
		pgx.NamedArgs{
			"id":       id,
			"username": "username",
		},
	)

	if err := row.Scan(&u.Id, &u.Username); err != nil {
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
