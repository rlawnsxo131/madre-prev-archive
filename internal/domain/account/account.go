package account

import (
	"time"
)

type Account struct {
	user          *User
	socialAccount *SocialAccount
}

type PublicAccount struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	PhotoUrl  *string   `json:"photo_url"`
	SocialID  string    `json:"social_id"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ac *Account) User() *User {
	return ac.user
}

func (ac *Account) SetUser(u *User) {
	ac.user = u
}

func (ac *Account) SocialAccount() *SocialAccount {
	return ac.socialAccount
}

func (ac *Account) SetSocialAccount(sa *SocialAccount) {
	ac.socialAccount = sa
}

func (ac *Account) PublicAccount() *PublicAccount {
	pa := PublicAccount{
		UserID:   ac.user.ID,
		Username: ac.user.Username,
		Email:    ac.user.Email,
		SocialID: ac.socialAccount.SocialID,
		Provider: ac.socialAccount.Provider,
	}

	if ac.user.PhotoUrl.Valid {
		pa.PhotoUrl = &ac.user.PhotoUrl.String
	} else {
		pa.PhotoUrl = nil
	}

	return &pa
}
