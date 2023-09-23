package user

type UserRepository interface {
	CreateForSocial(u *User) (string, error)
}

type UserQueryRepository interface {
	FindById(id string) (*User, error)
	FindByUsername(username string) (*User, error)
	ExistsByUsername(username string) (bool, error)
}

type UserSocialAccountQueryRepository interface {
	FindByUserId(userId string) (*UserSocialAccount, error)
	ExistsBySocialIdAndProvider(socialId, provider string) (bool, error)
}
