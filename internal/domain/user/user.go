package user

import (
	"regexp"
	"time"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/common"
	"github.com/rlawnsxo131/madre-server-v3/lib/logger"
	"github.com/rlawnsxo131/madre-server-v3/utils"
)

type User struct {
	Id            string         `json:"id"`
	Email         string         `json:"email"`
	Username      string         `json:"username"`
	PhotoUrl      string         `json:"photoUrl,omitempty"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	SocialAccount *SocialAccount `json:"socialAccount,omitempty"`
}

// public static constructor
func NewUserWithoutId(email, username, photoUrl string) (*User, error) {
	if email == "" || username == "" {
		return nil, common.ErrMissingRequiredValue
	}
	if err := validateUsername(username); err != nil {
		return nil, err
	}
	u := User{
		Email:    email,
		Username: username,
		PhotoUrl: photoUrl,
	}
	return &u, nil
}

func NewUserWithId(id, email, username, photoUrl string) (*User, error) {
	if id == "" || email == "" || username == "" {
		return nil, common.ErrMissingRequiredValue
	}
	if err := validateUsername(username); err != nil {
		return nil, err
	}
	u := User{
		Id:       id,
		Email:    email,
		Username: username,
		PhotoUrl: photoUrl,
	}
	return &u, nil
}

// public method
func (u *User) SetNewSocialAccount(socialId, socialUsername, provider string) error {
	if socialId == "" || socialUsername == "" || provider == "" {
		return common.ErrMissingRequiredValue
	}
	if err := u.isSupportSocialProvider(provider); err != nil {
		return err
	}
	sa := SocialAccount{
		SocialId:       socialId,
		SocialUsername: socialUsername,
		Provider:       provider,
	}
	u.SocialAccount = &sa
	return nil
}

func (u *User) SetSocialAccount(sa *SocialAccount) {
	u.SocialAccount = sa
}

// private static
func validateUsername(username string) error {
	if match, err := regexp.MatchString("^[a-zA-Z0-9]{1,20}$", username); err != nil {
		return errors.Wrap(err, "username regex MatchString error")
	} else if !match {
		logger.DefaultLogger().Error().Timestamp().
			Str("username", username).Msg("username regex validation failed")
		return common.ErrUnProcessableValue
	}
	return nil
}

func (u *User) isSupportSocialProvider(provider string) error {
	if isContain := utils.Contains([]string{SOCIAL_PROVIDER_GOOGLE}, provider); !isContain {
		logger.DefaultLogger().Error().Timestamp().
			Str("provider", provider).Msg("not support provider")
		return common.ErrNotSupportValue
	}
	return nil
}
