package user_test

import (
	"testing"

	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/domain/user"
	"github.com/rlawnsxo131/madre-server-v2/utils"
	"github.com/stretchr/testify/assert"
)

func Test_UserWriteUseCase_Create_IsSuccess(t *testing.T) {
	assert := assert.New(t)

	db, _ := database.DatabaseInstance()

	u := user.User{
		Email:      "madre@gmail.com",
		OriginName: utils.NewNullString("madre"),
		Username:   "madre",
		PhotoUrl:   utils.NewNullString("https://google.com"),
	}

	userWriteUseCase := user.NewWriteUseCase(db)
	id, err := userWriteUseCase.Create(&u)

	assert.Nil(err)
	assert.NotEmpty(id)
}

func Test_UserWriteUseCase_Create_IsFail(t *testing.T) {
	assert := assert.New(t)

	db, _ := database.DatabaseInstance()

	u := user.User{
		OriginName: utils.NewNullString("madre"),
		Username:   "madre",
		PhotoUrl:   utils.NewNullString("https://google.com"),
	}

	userWriteUseCase := user.NewWriteUseCase(db)
	id, err := userWriteUseCase.Create(&u)

	assert.Error(err)
	assert.Empty(id)
}
