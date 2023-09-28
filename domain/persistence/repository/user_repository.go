package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/huandu/go-sqlbuilder"
	"github.com/rlawnsxo131/madre-server/domain/entity/user"
	"github.com/rlawnsxo131/madre-server/domain/persistence"
	"github.com/rlawnsxo131/madre-server/domain/persistence/model"
)

type UserRepository struct {
	l      *persistence.QueryLogger
	mapper model.UserMapper
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		l:      persistence.NewQueryLogger(),
		mapper: model.UserMapper{},
	}
}

var _userStruct = sqlbuilder.NewStruct(&model.User{})

func (ur *UserRepository) FindById(
	ctx context.Context,
	id int64,
	opts *persistence.QueryOptions,
) (*user.User, error) {
	sb := _userStruct.SelectFrom("user")
	sb.Where(sb.Equal("id", id))

	if opts.WithTx {
		sb.ForUpdate()
	}

	query, args := sb.Build()

	ur.l.Logging(query, args)

	var u model.User
	err := opts.DB.
		QueryRowContext(ctx, query, args...).
		Scan(_userStruct.Addr(&u)...)

	if err != nil {
		return nil, err
	}

	return ur.mapper.MapToEntity(&u), nil
}

func (ur *UserRepository) ExistsUsername(
	ctx context.Context,
	username string,
	opts *persistence.QueryOptions,
) (bool, error) {
	sb := sqlbuilder.NewSelectBuilder()
	existsSb := sqlbuilder.NewSelectBuilder()

	query, args := sb.
		Select("true").
		From("user").
		Where(
			sb.Exists(
				existsSb.
					Select("1").
					From("user").
					Where(
						existsSb.Equal("username", username),
					),
			),
		).Build()

	ur.l.Logging(query, args)

	var exists bool
	err := opts.DB.
		QueryRowContext(ctx, query, args...).
		Scan(&exists)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return exists, nil
}

// List select
// sb := _userStruct.SelectFrom("user")
// sb.Where(sb.Equal("id", id))

// sql, args := sb.Build()

// var users []*user.User
// rows, err := db.Query(sql, args...)
// if err != nil {
// 	return nil, err
// }
// defer rows.Close()
// for rows.Next() {
// 	var u model.User
// 	rows.Scan(_userStruct.Addr(&u)...)
// 	fmt.Printf("current: %+v", ur.mapper.MapToEntity(&u))
// 	users = append(users, ur.mapper.MapToEntity(&u))
// }

// fmt.Println("len: ", len(users))
// fmt.Printf("users: %+v\n", &users[0])
// fmt.Println("len: ", len(users))
