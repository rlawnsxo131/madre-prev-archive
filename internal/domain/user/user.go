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
	PhotoUrl      string         `json:"photo_url,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	SocialAccount *SocialAccount `json:"social_account,omitempty"`
}

func NewSignUpUser(email, username, photoUrl, socialId, socialUsername, provider string) (*User, error) {
	if email == "" || username == "" || photoUrl == "" || socialId == "" || socialUsername == "" {
		return nil, common.ErrMissingRequiredValue
	}
	if err := validateUsername(username); err != nil {
		return nil, err
	}
	if err := isSupportSocialProvider(provider); err != nil {
		return nil, err
	}

	sa := SocialAccount{
		SocialId:       socialId,
		SocialUsername: socialUsername,
		Provider:       provider,
	}
	u := User{
		Email:         email,
		Username:      username,
		PhotoUrl:      photoUrl,
		SocialAccount: &sa,
	}
	return &u, nil
}

func (u *User) SetSocialAccount(sa *SocialAccount) {
	u.SocialAccount = sa
}

// private
func validateUsername(username string) error {
	if match, err := regexp.MatchString("^[a-zA-Z0-9]{1,20}$", username); err != nil {
		return errors.Wrap(err, "username regex MatchString error")
	} else if !match {
		logger.DefaultLogger().Error().Timestamp().
			Str("username", username).Msgf("username regex validation failed")
		return common.ErrUnProcessableValue
	}
	return nil
}

func isSupportSocialProvider(provider string) error {
	if isContain := utils.Contains([]string{SOCIAL_PROVIDER_GOOGLE}, provider); !isContain {
		logger.DefaultLogger().Error().Timestamp().
			Str("provider", provider).Msg("not support provider")
		return common.ErrNotSupportValue
	}
	return nil
}
