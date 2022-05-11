package user_test

import (
	"testing"

	"github.com/rlawnsxo131/madre-server-v2/domain/user"
	"github.com/stretchr/testify/assert"
)

func Test_User_ValidateUsername_ValidIsFalse(t *testing.T) {
	assert := assert.New(t)

	u := &user.User{}
	valid, _ := u.ValidateUsername()

	assert.False(valid)
}

func Test_User_ValidateUsername_ValidIsTrue(t *testing.T) {
	assert := assert.New(t)

	u := &user.User{
		Username: "username",
	}
	valid, _ := u.ValidateUsername()

	assert.True(valid)
}
