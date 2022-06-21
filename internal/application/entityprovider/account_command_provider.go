package entityprovider

import (
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/account"
	"github.com/rlawnsxo131/madre-server-v3/utils"
)

func NewSaveAccountUser(email, originName, username, photoUrl string) *account.User {
	return &account.User{
		Email:      email,
		OriginName: utils.NewNullString(originName),
		Username:   username,
		PhotoUrl:   utils.NewNullString(photoUrl),
	}
}

func NewSaveAccountSocialAccount(socialId, provider string) *account.SocialAccount {
	return &account.SocialAccount{
		SocialID: socialId,
		Provider: provider,
	}
}
