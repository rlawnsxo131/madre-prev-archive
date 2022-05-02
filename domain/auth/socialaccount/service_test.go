package socialaccount_test

import (
	"testing"

	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/domain/auth/socialaccount"
	"github.com/rlawnsxo131/madre-server-v2/utils"
	"github.com/stretchr/testify/assert"
)

func Test_Service_Create_IsSuccess(t *testing.T) {
	assert := assert.New(t)

	db, _ := database.GetDatabaseInstance()

	socialAccount := socialaccount.SocialAccount{
		UserID:   utils.GenerateUUIDString(),
		Provider: "GOOGLE",
		SocialId: utils.GenerateUUIDString(),
	}

	socialAccountService := socialaccount.NewService(db)
	id, err := socialAccountService.Create(socialAccount)

	assert.Nil(err)
	assert.NotEmpty(id)
}

func Test_Service_Create_IsFail(t *testing.T) {
	assert := assert.New(t)

	db, _ := database.GetDatabaseInstance()

	socialAccount := socialaccount.SocialAccount{
		UserID:   utils.GenerateUUIDString(),
		Provider: "",
		SocialId: utils.GenerateUUIDString(),
	}

	socialAccountService := socialaccount.NewService(db)
	id, err := socialAccountService.Create(socialAccount)

	assert.Error(err)
	assert.Empty(id)
}
