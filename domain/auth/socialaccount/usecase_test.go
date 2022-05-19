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

	db, _ := database.DatabaseInstance()

	sa := socialaccount.SocialAccount{
		UserID:   uuid.NewString(),
		SocialId: uuid.NewString(),
		Provider: "GOOGLE",
	}

	socialUseCase := socialaccount.NewUseCase(db)
	id, err := socialUseCase.Create(&sa)

	assert.Nil(err)
	assert.NotEmpty(id)
}

func Test_SocialAccountUseCase_Create_IsFail(t *testing.T) {
	assert := assert.New(t)

	db, _ := database.DatabaseInstance()

	sa := socialaccount.SocialAccount{
		UserID:   uuid.NewString(),
		SocialId: uuid.NewString(),
		Provider: "",
	}

	socialUseCase := socialaccount.NewUseCase(db)
	id, err := socialUseCase.Create(&sa)

	assert.Error(err)
	assert.Empty(id)
}

func Test_SocialAccountUseCase_FindOneByProviderWithSocialId_IsSuccess(t *testing.T) {
	assert := assert.New(t)

	db, _ := database.DatabaseInstance()

	socialId := uuid.NewString()
	sa := socialaccount.SocialAccount{
		UserID:   uuid.NewString(),
		SocialId: socialId,
		Provider: socialaccount.Key_Provider_GOOGLE,
	}

	socialUseCase := socialaccount.NewUseCase(db)
	id, _ := socialUseCase.Create(&sa)

	newSa, err := socialUseCase.FindOneBySocialIdWithProvider(
		socialId,
		socialaccount.Key_Provider_GOOGLE,
	)

	assert.Nil(err)
	assert.Equal(id, newSa.ID)
}
