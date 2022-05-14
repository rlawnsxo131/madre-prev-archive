package socialaccount_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/domain/auth/socialaccount"
	"github.com/stretchr/testify/assert"
)

func Test_SoicalAccountUseCase_Create_IsSuccess(t *testing.T) {
	assert := assert.New(t)

	db, _ := database.GetDatabaseInstance()

	sa := socialaccount.SocialAccount{
		UserID:   uuid.NewString(),
		Provider: "GOOGLE",
		SocialId: uuid.NewString(),
	}

	socialUseCase := socialaccount.NewUseCase(db)
	id, err := socialUseCase.Create(&sa)

	assert.Nil(err)
	assert.NotEmpty(id)
}

func Test_SocialAccountUseCase_Create_IsFail(t *testing.T) {
	assert := assert.New(t)

	db, _ := database.GetDatabaseInstance()

	sa := socialaccount.SocialAccount{
		UserID:   uuid.NewString(),
		Provider: "",
		SocialId: uuid.NewString(),
	}

	socialUseCase := socialaccount.NewUseCase(db)
	id, err := socialUseCase.Create(&sa)

	assert.Error(err)
	assert.Empty(id)
}
