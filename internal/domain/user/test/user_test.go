package user_test

import (
	"testing"

	"github.com/rlawnsxo131/madre-server-v3/internal/domain/common"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
	"github.com/stretchr/testify/assert"
)

func Test_SetNewUserWithoutId_Success(t *testing.T) {
	assert := assert.New(t)

	email := "email"
	username := "username"
	photoUrl := "photoUrl"

	u, err := user.NewUserWithoutId(
		email,
		username,
		photoUrl,
	)

	assert.Nil(err)
	assert.Equal(email, u.Email)
	assert.Equal(username, u.Username)
	assert.Equal(photoUrl, u.PhotoUrl)
}

func Test_SetNewUserWithoutId_Return_ErrMissingRequiredValue(t *testing.T) {
	assert := assert.New(t)

	email := ""
	username := "username"
	photoUrl := "photoUrl"

	u, err := user.NewUserWithoutId(
		email,
		username,
		photoUrl,
	)

	assert.Nil(u)
	assert.ErrorIs(err, common.ErrMissingRequiredValue)
}

func Test_SetSocialAccount_Success(t *testing.T) {
	assert := assert.New(t)

	u := &user.User{}

	socialId := "socialId"
	socialUsername := "socialUsername"
	provider := user.SOCIAL_PROVIDER_GOOGLE

	err := u.SetNewSocialAccount(
		socialId,
		socialUsername,
		provider,
	)

	assert.Nil(err)
}

func Test_SetSocialAccount_Return_ErrNotSupportValue(t *testing.T) {
	assert := assert.New(t)

	u := &user.User{}

	socialId := "socialId"
	socialUsername := "socialUsername"
	provider := "not supported value"

	err := u.SetNewSocialAccount(
		socialId,
		socialUsername,
		provider,
	)

	assert.ErrorIs(err, common.ErrNotSupportValue)
}
