package user

type UserUseCase interface {
	Create(u *User) (string, error)
	FindOneById(id string) (*User, error)
	FindOneByUsername(username string) (*User, error)
}

type userUseCase struct {
	repo UserRepository
}

func NewUserUseCase(repo UserRepository) UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (uc *userUseCase) Create(u *User) (string, error) {
	return uc.repo.Create(u)
}

func (uc *userUseCase) FindOneById(id string) (*User, error) {
	return uc.repo.FindOneById(id)
}

func (uc *userUseCase) FindOneByUsername(username string) (*User, error) {
	return uc.repo.FindOneByUsername(username)
}
