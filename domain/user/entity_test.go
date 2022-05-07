package user_test

import (
	"testing"

	"github.com/rlawnsxo131/madre-server-v2/domain/user"
	"github.com/stretchr/testify/assert"
)

func Test_User_ValidateDisplayName_ValidIsFalse(t *testing.T) {
	assert := assert.New(t)

	u := &user.User{}
	valid, _ := u.ValidateDisplayName()

	assert.False(valid)
}

func Test_User_ValidateDisplayName_ValidIsTrue(t *testing.T) {
	assert := assert.New(t)

	u := &user.User{
		DisplayName: "displayName",
	}
	valid, _ := u.ValidateDisplayName()

	assert.True(valid)
}
