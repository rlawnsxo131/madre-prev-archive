package user_test

import (
	"testing"

	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/domain/user"
	"github.com/rlawnsxo131/madre-server-v2/utils"
	"github.com/stretchr/testify/assert"
)

func Test_UserUseCase_Create_IsSuccess(t *testing.T) {
	assert := assert.New(t)

	db, _ := database.GetDatabaseInstance()

	u := user.User{
		Email:      "madre@gmail.com",
		OriginName: utils.NewNullString("madre"),
		Username:   "madre",
		PhotoUrl:   utils.NewNullString("https://google.com"),
	}

	userUseCase := user.NewUseCase(db)
	id, err := userUseCase.Create(&u)

	assert.Nil(err)
	assert.NotEmpty(id)
}

func Test_UserUseCase_Create_IsFail(t *testing.T) {
	assert := assert.New(t)

	db, _ := database.GetDatabaseInstance()

	u := user.User{
		OriginName: utils.NewNullString("madre"),
		Username:   "madre",
		PhotoUrl:   utils.NewNullString("https://google.com"),
	}

	userUseCase := user.NewUseCase(db)
	id, err := userUseCase.Create(&u)

	assert.Error(err)
	assert.Empty(id)
}

func Test_UserUseCase_FindOneById_IsSuccess(t *testing.T) {}

func Test_UserUseCase_FindOneById_IsFail(t *testing.T) {}

func Test_USerUseCase_FindOneByUsername_IsSuccess(t *testing.T) {}

func Test_USerUseCase_FindOneByUsername_IsFail(t *testing.T) {}
