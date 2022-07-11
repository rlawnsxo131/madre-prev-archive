package account

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

var (
	ErrInvalidUsername       = errors.New("username is required")
	ErrInvalidEmail          = errors.New("email is required")
	ErrInvalidSocialId       = errors.New("socialId is required")
	ErrInvalidProvider       = errors.New("provider is required")
	ErrInvalidProviderGoogle = errors.New("provider is not google")
)

type Account struct {
	user          *User
	socialAccount *SocialAccount
}

type PublicAccount struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	PhotoUrl  *string   `json:"photo_url"`
	SocialID  string    `json:"social_id"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAccount(
	username string,
	email string,
	photoUrl sql.NullString,
	socialId string,
	provider string,
) (*Account, error) {
	if username == "" {
		return nil, ErrInvalidUsername
	}
	if email == "" {
		return nil, ErrInvalidEmail
	}
	if socialId == "" {
		return nil, ErrInvalidSocialId
	}
	if provider == "" {
		return nil, ErrInvalidProvider
	}
	if provider != SOCIAL_PROVIDER_GOOGLE {
		return nil, ErrInvalidProviderGoogle
	}

	u := &User{
		Username: username,
		Email:    email,
		PhotoUrl: photoUrl,
	}
	err := u.ValidateUsername()
	if err != nil {
		return nil, err
	}

	sa := &SocialAccount{
		SocialID: socialId,
		Provider: provider,
	}

	ac := &Account{u, sa}
	return ac, nil
}

func (ac *Account) User() *User {
	return ac.user
}

func (ac *Account) AddUser(u *User) {
	ac.user = u
}

func (ac *Account) SocialAccount() *SocialAccount {
	return ac.socialAccount
}

func (ac *Account) AddSocialAccount(sa *SocialAccount) {
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
