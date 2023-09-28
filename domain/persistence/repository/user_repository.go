package repository

import (
	"context"

	"github.com/huandu/go-sqlbuilder"
	"github.com/rlawnsxo131/madre-server/domain/entity/user"
	"github.com/rlawnsxo131/madre-server/domain/persistence"
	"github.com/rlawnsxo131/madre-server/domain/persistence/model"
)

type userRepository struct {
	l      *persistence.QueryLogger
	mapper model.UserMapper
}

func NewUserRepository() *userRepository {
	return &userRepository{
		l:      persistence.NewQueryLogger(),
		mapper: model.UserMapper{},
	}
}

var _userStruct = sqlbuilder.NewStruct(&model.User{})

func (ur *userRepository) FindById(ctx context.Context, id int64, opts *persistence.QueryOptions) (*user.User, error) {
	sb := _userStruct.SelectFrom("user")
	sb.Where(sb.Equal("id", id))

	if opts != nil && opts.WithTx {
		sb.ForUpdate()
	}

	sql, args := sb.Build()
	ur.l.Logging(sql, args)

	var u model.User
	err := opts.DB.
		QueryRowContext(ctx, sql, args...).
		Scan(_userStruct.Addr(&u)...)

	if err != nil {
		return nil, err
	}

	return ur.mapper.MapToEntity(&u), nil
}

// use
// repo := repository.NewUserRepository()
// u, err := repo.FindById(
// 	context.Background(), 1,
// 	&persistence.QueryOptions{
// 		DB:     db,
// 		WithTx: true,
// 	},
// )

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
