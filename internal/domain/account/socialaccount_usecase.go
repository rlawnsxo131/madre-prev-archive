package account

type SocialAccountUseCase interface {
	Create(sa *SocialAccount) (string, error)
	FindOneBySocialIdAndProvider(socialId, provider string) (*SocialAccount, error)
}
