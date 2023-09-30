package user

import (
	"strings"
	"time"

	"github.com/rlawnsxo131/madre-server/core/adapter"
	"github.com/rlawnsxo131/madre-server/domain/errz"
)

// @TODO 테스트 재작성

type User struct {
	Id            int64              `json:"id"`
	Email         string             `json:"email"`
	Username      string             `json:"username"`
	PhotoUrl      string             `json:"photoUrl,omitempty"`
	CreatedAt     time.Time          `json:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt"`
	SocialAccount *userSocialAccount `json:"socialAccount,omitempty"`
}

func New(email, photoUrl string) (*User, error) {
	var params = struct {
		Email    string `validate:"required,email"`
		PhotoUrl string `validate:"omitempty,url"`
	}{
		Email:    email,
		PhotoUrl: photoUrl,
	}

	if err := adapter.Validator().Struct(params); err != nil {
		return nil, err
	}

	// initial name is generated as uuid
	return &User{
		Email:    email,
		Username: strings.ReplaceAll("uuid 만들어야 함", "-", ""),
		PhotoUrl: photoUrl,
	}, nil
}

func (u *User) SetUsername(username string) error {
	var params = struct {
		Username string `validate:"required,alphanum,max=50"`
	}{
		Username: username,
	}

	if err := adapter.Validator().Struct(params); err != nil {
		return err
	}

	u.Username = username

	return nil
}

func (u *User) SetNewSocialAccount(socialId, provider string) error {
	socialAccount, err := newUserSocialAccount(u.Id, socialId, provider)

	if err != nil {
		return err
	}

	u.SocialAccount = socialAccount

	return nil
}

func (u *User) SetSocialAccount(sa *userSocialAccount) error {
	if sa == nil {
		return errz.NewErrMissingRequiredValue(sa)
	}

	u.SocialAccount = sa

	return nil
}
