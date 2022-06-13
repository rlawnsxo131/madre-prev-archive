package application

import "github.com/rlawnsxo131/madre-server-v3/internal/domain/account"

type userUseCase struct {
	commandRepository account.UserCommandRepository
	queryRepository   account.UserQueryRepository
}

func NewUserUseCase(
	commandRepository account.UserCommandRepository,
	queryRepository account.UserQueryRepository,
) account.UserUseCase {
	return &userUseCase{
		commandRepository,
		queryRepository,
	}
}

func (s *userUseCase) Create(u *account.User) (string, error) {
	return s.commandRepository.Create(u)
}

func (s *userUseCase) FindOneById(id string) (*account.User, error) {
	return s.queryRepository.FindOneById(id)
}

func (s *userUseCase) FindOneByUsername(username string) (*account.User, error) {
	return s.queryRepository.FindOneByUsername(username)
}
