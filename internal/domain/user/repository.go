package user

type UserRepository interface {
	Create(u *User) (string, error)
}

type UserQueryRepository interface {
	FindById(id string) (*User, error)
	FindByUsername(username string) (*User, error)
	ExistsByUsername(username string) (bool, error)
}

type SocialAccountQueryRepository interface {
	FindByUserId(userId string) (*SocialAccount, error)
	FindBySocialIdAndProvider(socialId, provider string) (*SocialAccount, error)
	ExistsBySocialIdAndProvider(socialId, provider string) (bool, error)
}
