package infrastructure

import "github.com/rlawnsxo131/madre-server-v3/internal/domain/account"

type AccountMapper struct{}

func (am AccountMapper) ToUserEntity(u *account.User) *account.User {
	return &account.User{
		ID:         u.ID,
		Email:      u.Email,
		OriginName: u.OriginName,
		Username:   u.Username,
		PhotoUrl:   u.PhotoUrl,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

func (am AccountMapper) ToUserModel(u *account.User) *account.User {
	return &account.User{
		Email:      u.Email,
		OriginName: u.OriginName,
		Username:   u.Username,
		PhotoUrl:   u.PhotoUrl,
		UpdatedAt:  u.UpdatedAt,
	}
}

func (am AccountMapper) ToSocialAccountEntity(sa *account.SocialAccount) *account.SocialAccount {
	return &account.SocialAccount{
		ID:        sa.ID,
		UserID:    sa.UserID,
		Provider:  sa.Provider,
		SocialID:  sa.SocialID,
		CreatedAt: sa.CreatedAt,
		UpdatedAt: sa.UpdatedAt,
	}
}

func (am AccountMapper) ToSocialAccountModel(sa *account.SocialAccount) *account.SocialAccount {
	return &account.SocialAccount{
		UserID:   sa.UserID,
		Provider: sa.Provider,
		SocialID: sa.SocialID,
	}
}
