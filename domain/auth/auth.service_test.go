package auth_test

import (
	"database/sql"
	"testing"

	"github.com/rlawnsxo131/madre-server-v2/domain/auth"
	"github.com/stretchr/testify/assert"
)

func Test_AuthService_GetExistSocialAccountMap_ExistIsFalse(t *testing.T) {
	assert := assert.New(t)

	err := sql.ErrNoRows
	socialAccount := auth.SocialAccount{
		ID:   0,
		UUID: "uuid",
	}

	authService := auth.NewAuthService()
	existSocialAccountMap, err := authService.GetExistSocialAccountMap(socialAccount, err)

	assert.Equal(err, nil)
	assert.False(existSocialAccountMap["exist"])
}

func Test_AuthService_GetExistSocialAccountMap_ExistIsTrue(t *testing.T) {
	assert := assert.New(t)

	socialAccount := auth.SocialAccount{
		ID:   1,
		UUID: "uuid",
	}

	authService := auth.NewAuthService()
	existSocialAccountMap, err := authService.GetExistSocialAccountMap(socialAccount, nil)

	assert.Equal(err, nil)
	assert.True(existSocialAccountMap["exist"])
}

func Test_AuthService_ValidateDisplayName_ValidIsFalse(t *testing.T) {
	assert := assert.New(t)

	displayName := ""

	authService := auth.NewAuthService()
	valid, _ := authService.ValidateDisplayName(displayName)

	assert.False(valid)
}

func Test_AuthService_ValidateDisplayName_ValidIsTrue(t *testing.T) {
	assert := assert.New(t)

	displayName := "displayName"

	authService := auth.NewAuthService()
	valid, _ := authService.ValidateDisplayName(displayName)

	assert.True(valid)
}
