package queryrepository

import "github.com/rlawnsxo131/madre-server-v3/internal/domain/user"

type userQueryRepository struct{}

func NewUserQueryRepository() user.UserQueryRepository {
	return &userQueryRepository{}
}

func (uqr *userQueryRepository) FindById(id string) (*user.User, error)
func (uqr *userQueryRepository) FindByUsername(username string) (*user.User, error)
func (uqr *userQueryRepository) ExistsByUsername(username string) (bool, error)
