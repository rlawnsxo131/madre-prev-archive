package user

import (
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

type UserRepository interface {
	Create(u *User) (string, error)
	FindOneById(id string) (*User, error)
	FindOneByUsername(username string) (*User, error)
}

type userRepository struct {
	db     database.Database
	mapper userEntityMapper
}

func NewUserRepository(db database.Database) UserRepository {
	return &userRepository{
		db:     db,
		mapper: userEntityMapper{},
	}
}

func (r *userRepository) Create(u *User) (string, error) {
	var id string

	query := "INSERT INTO public.user(email, origin_name, username, photo_url)" +
		" VALUES(:email, :origin_name, :username, :photo_url)" +
		" RETURNING id"

	err := r.db.PrepareNamedGet(
		&id,
		query,
		r.mapper.toModel(u),
	)
	if err != nil {
		return "", errors.Wrap(err, "user WriteRepository create")
	}

	return id, nil
}

func (r *userRepository) FindOneById(id string) (*User, error) {
	var u User

	query := "SELECT * FROM public.user" +
		" WHERE id = $1"

	err := r.db.QueryRowx(query, id).StructScan(&u)
	if err != nil {
		customError := errors.Wrap(err, "user ReadRepository FindOneById")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return r.mapper.toEntity(&u), err
}

func (r *userRepository) FindOneByUsername(username string) (*User, error) {
	var u User

	query := "SELECT * FROM public.user" +
		" WHERE username = $1"

	err := r.db.QueryRowx(query, username).StructScan(&u)
	if err != nil {
		customError := errors.Wrap(err, "user ReadRepository FindOneByUsername")
		err = utils.ErrNoRowsReturnRawError(err, customError)
	}

	return r.mapper.toEntity(&u), err
}
