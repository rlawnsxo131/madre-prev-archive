package command

import (
	"github.com/rlawnsxo131/madre-server/domain/entity/user"
	"github.com/rlawnsxo131/madre-server/domain/persistence"
	"github.com/rlawnsxo131/madre-server/domain/persistence/repository"
)

type UserCommandService struct {
	conn     persistence.Conn
	userRepo *repository.UserRepository
}

func NewUserCommandService(conn persistence.Conn) *UserCommandService {
	return &UserCommandService{
		conn:     conn,
		userRepo: repository.NewUserRepository(),
	}
}

func (ucs *UserCommandService) Create(command CreateUserCommand) (*user.User, error) {
	return nil, nil
}
