package auth_test

import (
	"database/sql"
	"testing"

	"github.com/rlawnsxo131/madre-server-v2/domain/auth"
	"github.com/rlawnsxo131/madre-server-v2/domain/auth/socialaccount"
	"github.com/rlawnsxo131/madre-server-v2/utils"
	"github.com/stretchr/testify/assert"
)

func Test_AuthService_GetExistSocialAccountMap_ExistIsFalse(t *testing.T) {
	assert := assert.New(t)

	err := sql.ErrNoRows
	socialAccount := &socialaccount.SocialAccount{
		ID: "",
	}

	authService := auth.NewService()
	existSocialAccountMap, err := authService.GetExistSocialAccountMap(socialAccount, err)

	assert.Nil(err)
	assert.False(existSocialAccountMap["exist"])
}

func Test_AuthService_GetExistSocialAccountMap_ExistIsTrue(t *testing.T) {
	assert := assert.New(t)

	socialAccount := &socialaccount.SocialAccount{
		ID: utils.GenerateUUIDString(),
	}

	authService := auth.NewService()
	existSocialAccountMap, err := authService.GetExistSocialAccountMap(socialAccount, nil)

	assert.Nil(err)
	assert.True(existSocialAccountMap["exist"])
}

func Test_AuthService_ValidateDisplayName_ValidIsFalse(t *testing.T) {
	assert := assert.New(t)

	displayName := ""

	authService := auth.NewService()
	valid, _ := authService.ValidateDisplayName(displayName)

	assert.False(valid)
}

func Test_AuthService_ValidateDisplayName_ValidIsTrue(t *testing.T) {
	assert := assert.New(t)

	displayName := "displayName"

	authService := auth.NewService()
	valid, _ := authService.ValidateDisplayName(displayName)

	assert.True(valid)
}
