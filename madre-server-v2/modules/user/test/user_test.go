package user_test

import (
	"testing"

	"github.com/rlawnsxo131/madre-server-v2/modules/user"
	"github.com/stretchr/testify/assert"
)

func Test_User_Filter(t *testing.T) {}

func Test_User_IsExist_ExistIsTrue(t *testing.T) {}

func Test_User_IsExist_ExistIsFalse(t *testing.T) {}

func Test_User_ValidateUsername_ValidIsTrue(t *testing.T) {
	assert := assert.New(t)

	u := &user.User{
		Username: "username",
	}
	valid, _ := u.ValidateUsername()

	assert.True(valid)
}

func Test_User_ValidateUsername_ValidIsFalse(t *testing.T) {
	assert := assert.New(t)

	u := &user.User{}
	valid, _ := u.ValidateUsername()

	assert.False(valid)
}
