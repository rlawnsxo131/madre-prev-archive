package query

import (
	"context"
	"net/http"

	"github.com/rlawnsxo131/madre-server/api/service"
	"github.com/rlawnsxo131/madre-server/core/errorz"
	"github.com/rlawnsxo131/madre-server/domain/errz"
	"github.com/rlawnsxo131/madre-server/domain/persistence"
	"github.com/rlawnsxo131/madre-server/domain/persistence/repository"
)

type UserQueryService struct {
	conn     persistence.Conn
	userRepo *repository.UserRepository
}

func NewUserQueryService(conn persistence.Conn) *UserQueryService {
	return &UserQueryService{
		conn:     conn,
		userRepo: repository.NewUserRepository(),
	}
}

func (uqs *UserQueryService) IsExistsUsername(query IsExistsUsernameQuery) *service.ErrWithHTTPCode {
	exists, err := uqs.userRepo.ExistsByUsername(
		context.Background(),
		&persistence.QueryOptions{
			Conn: uqs.conn,
		},
		query.Username,
	)

	switch {
	case err != nil:
		return service.NewErrWithHTTPCode(
			errorz.New(err),
			http.StatusConflict,
			"시스템 에러",
		)
	case exists:
		return service.NewErrWithHTTPCode(
			errorz.New(
				errz.NewErrConflictUniqValue(query.Username),
			),
			http.StatusConflict,
			"중복된 이름입니다",
		)
	}

	return nil
}
