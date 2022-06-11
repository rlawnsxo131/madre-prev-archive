package user

type UserService interface {
	Create(u *User) (string, error)
	FindOneById(id string) (*User, error)
	FindOneByUsername(username string) (*User, error)
}

type userService struct {
	commandRepo UserCommandRepository
	queryRepo   UserQueryRepository
}

func NewUserService(
	commandRepo UserCommandRepository,
	queryRepo UserQueryRepository,
) UserService {
	return &userService{
		commandRepo,
		queryRepo,
	}
}

func (s *userService) Create(u *User) (string, error) {
	return s.commandRepo.Create(u)
}

func (s *userService) FindOneById(id string) (*User, error) {
	return s.queryRepo.FindOneById(id)
}

func (s *userService) FindOneByUsername(username string) (*User, error) {
	return s.queryRepo.FindOneByUsername(username)
}
