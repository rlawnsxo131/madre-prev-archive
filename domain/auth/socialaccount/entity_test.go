package socialaccount_test

import (
	"testing"

	"github.com/rlawnsxo131/madre-server-v2/domain/auth/socialaccount"
	"github.com/rlawnsxo131/madre-server-v2/utils"
	"github.com/stretchr/testify/assert"
)

func Test_SocialAccount_IsExist_IsTrue(t *testing.T) {
	assert := assert.New(t)

	sa := &socialaccount.SocialAccount{
		ID: utils.GenerateUUIDString(),
	}
	exist, err := sa.IsExist(nil)

	assert.Nil(err)
	assert.True(exist)
}
