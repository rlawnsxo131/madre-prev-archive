package socialaccount_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/domain/auth/socialaccount"
	"github.com/stretchr/testify/assert"
)

func Test_SocialAccountReadUseCase_FindOneByProviderWithSocialId_IsSuccess(t *testing.T) {
	assert := assert.New(t)

	db, _ := database.DatabaseInstance()

	socialId := uuid.NewString()
	sa := socialaccount.SocialAccount{
		UserID:   uuid.NewString(),
		SocialID: socialId,
		Provider: socialaccount.Key_Provider_GOOGLE,
	}

	socialReadUseCase := socialaccount.NewReadUseCase(db)
	socialWriteUseCase := socialaccount.NewWriteUseCase(db)
	id, _ := socialWriteUseCase.Create(&sa)

	newSa, err := socialReadUseCase.FindOneBySocialIdAndProvider(
		&socialaccount.SocialIDAndProviderDto{
			SocialID: sa.SocialID,
			Provider: socialaccount.Key_Provider_GOOGLE,
		},
	)

	assert.Nil(err)
	assert.Equal(id, newSa.ID)
}
