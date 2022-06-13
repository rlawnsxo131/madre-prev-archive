package account_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/account"
	"github.com/stretchr/testify/assert"
)

func Test_SocialAccount_IsExist_IsTrue(t *testing.T) {
	assert := assert.New(t)

	sa := &account.SocialAccount{
		ID: uuid.NewString(),
	}
	exist, err := sa.IsExist(nil)

	assert.Nil(err)
	assert.True(exist)
}
