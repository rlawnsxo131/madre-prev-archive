package user_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
	"github.com/stretchr/testify/assert"
)

func Test_User_IsExist_IsTrue(t *testing.T) {
	assert := assert.New(t)

	u := &user.User{
		ID: uuid.NewString(),
	}
	exist, err := u.IsExist(nil)

	assert.Nil(err)
	assert.True(exist)
}

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
