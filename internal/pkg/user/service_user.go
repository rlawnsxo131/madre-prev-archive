package user

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Create(u *User) (string, error) {
	return s.repo.Create(u)
}

func (s *userService) FindOneById(id string) (*User, error) {
	return s.repo.FindOneById(id)
}

func (s *userService) FindOneByUsername(username string) (*User, error) {
	return s.repo.FindOneByUsername(username)
}
