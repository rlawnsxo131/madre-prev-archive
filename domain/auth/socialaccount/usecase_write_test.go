package socialaccount_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/domain/auth/socialaccount"
	"github.com/stretchr/testify/assert"
)

func Test_SoicalAccountWriteUseCase_Create_IsSuccess(t *testing.T) {
	assert := assert.New(t)

	db, _ := database.DatabaseInstance()

	sa := socialaccount.SocialAccount{
		UserID:   uuid.NewString(),
		SocialId: uuid.NewString(),
		Provider: "GOOGLE",
	}

	socialWriteUseCase := socialaccount.NewWriteUseCase(db)
	id, err := socialWriteUseCase.Create(&sa)

	assert.Nil(err)
	assert.NotEmpty(id)
}

func Test_SoicalAccountWriteUseCase_Create_IsFail(t *testing.T) {
	assert := assert.New(t)

	db, _ := database.DatabaseInstance()

	sa := socialaccount.SocialAccount{
		UserID:   uuid.NewString(),
		SocialId: uuid.NewString(),
		Provider: "",
	}

	socialWriteUseCase := socialaccount.NewWriteUseCase(db)
	id, err := socialWriteUseCase.Create(&sa)

	assert.Error(err)
	assert.Empty(id)
}
