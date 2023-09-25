package repository

import (
	"github.com/huandu/go-sqlbuilder"
	"github.com/rlawnsxo131/madre-server/domain/entity/user"
	"github.com/rlawnsxo131/madre-server/domain/persistence"
	"github.com/rlawnsxo131/madre-server/domain/persistence/model"
)

type userRepository struct {
	mapper model.UserMapper
}

func NewUserRepository() *userRepository {
	return &userRepository{
		mapper: model.UserMapper{},
	}
}

var _userStruct = sqlbuilder.NewStruct(&model.User{})

func (ur *userRepository) FindById(queryLayer persistence.QueryLayer, id int64) (*user.User, error) {
	sb := _userStruct.SelectFrom("user")
	sb.Where(sb.Equal("id", id))

	sql, args := sb.Build()

	var u model.User

	err := queryLayer.
		QueryRow(sql, args).
		Scan(_userStruct.Addr(&u)...)

	if err != nil {
		return nil, err
	}

	return ur.mapper.MapToEntity(&u), nil
}
