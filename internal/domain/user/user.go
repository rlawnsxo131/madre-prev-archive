package user

import (
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/common"
	valueutil "github.com/rlawnsxo131/madre-server-v3/pkg/core/utils/value-util"
)

var (
	ErrUsernameRegexNotMatched = errors.Wrap(
		common.ErrUnprocessableValue,
		"username regex valdiation not matched",
	)
)

type User struct {
	Id            string             `json:"id"`
	Email         string             `json:"email"`
	Username      string             `json:"username"`
	PhotoUrl      string             `json:"photoUrl,omitempty"`
	CreatedAt     time.Time          `json:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt"`
	SocialAccount *UserSocialAccount `json:"socialAccount,omitempty"`
}

func NewUserWithoutId(email, photoUrl string) (*User, error) {
	if email == "" {
		return nil, common.ErrMissingRequiredValue
	}

	// initial name is generated as uuid
	return &User{
		Email:    email,
		Username: strings.ReplaceAll(valueutil.NewUUIDString(), "-", ""),
		PhotoUrl: photoUrl,
	}, nil
}

func NewUserWithId(id, username, email, photoUrl string) (*User, error) {
	if id == "" || username == "" || email == "" {
		return nil, common.ErrMissingRequiredValue
	}

	if err := validateUsername(username); err != nil {
		return nil, err
	}

	return &User{
		Id:       id,
		Username: username,
		Email:    email,
		PhotoUrl: photoUrl,
	}, nil
}

func (u *User) SetNewSocialAccount(socialId, socialUsername, provider string) error {
	if socialId == "" || provider == "" {
		return common.ErrMissingRequiredValue
	}
	if err := u.isSupportSocialProvider(provider); err != nil {
		return err
	}

	u.SocialAccount = &UserSocialAccount{
		SocialId:       socialId,
		SocialUsername: socialUsername,
		Provider:       provider,
	}

	return nil
}

func (u *User) SetSocialAccount(sa *UserSocialAccount) error {
	if sa == nil {
		return common.ErrMissingRequiredValue
	}
	u.SocialAccount = sa
	return nil
}

func (u *User) isSupportSocialProvider(provider string) error {
	isContain := valueutil.Contains(
		[]string{SOCIAL_PROVIDER_GOOGLE},
		provider,
	)
	if !isContain {
		return errors.Wrap(
			common.ErrNotSupportValue,
			"not support provider",
		)
	}
	return nil
}

func validateUsername(username string) error {
	match, err := regexp.MatchString(
		"^[a-zA-Z0-9]{1,20}$",
		username,
	)
	if err != nil {
		return errors.Wrap(
			err,
			"username regex MatchString parse error",
		)
	}
	if !match {
		return ErrUsernameRegexNotMatched
	}
	return nil
}
