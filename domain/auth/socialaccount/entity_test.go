package socialaccount_test

import (
	"testing"

	"github.com/rlawnsxo131/madre-server-v2/domain/auth/socialaccount"
	"github.com/rlawnsxo131/madre-server-v2/utils"
	"github.com/stretchr/testify/assert"
)

func Test_AuthService_GetExistSocialAccountMap_ExistIsTrue(t *testing.T) {
	assert := assert.New(t)

	sa := &socialaccount.SocialAccount{
		ID: utils.GenerateUUIDString(),
	}
	existSocialAccountMap, err := sa.GetExistSocialAccountMap(nil)

	assert.Nil(err)
	assert.True(existSocialAccountMap["exist"])
}
