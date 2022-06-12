package user

type UserService interface {
	Create(u *User) (string, error)
	FindOneById(id string) (*User, error)
	FindOneByUsername(username string) (*User, error)
}

type userService struct {
	command UserCommandRepository
	query   UserQueryRepository
}

func NewUserService(
	command UserCommandRepository,
	query UserQueryRepository,
) UserService {
	return &userService{
		command,
		query,
	}
}

func (s *userService) Create(u *User) (string, error) {
	return s.command.Create(u)
}

func (s *userService) FindOneById(id string) (*User, error) {
	return s.query.FindOneById(id)
}

func (s *userService) FindOneByUsername(username string) (*User, error) {
	return s.query.FindOneByUsername(username)
}
