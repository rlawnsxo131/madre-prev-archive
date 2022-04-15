package auth_test

import (
	"testing"

	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/domain/auth"
	"github.com/rlawnsxo131/madre-server-v2/utils"
	"github.com/stretchr/testify/assert"
)

func Test_SocialAccountService_Create_Success(t *testing.T) {
	assert := assert.New(t)

	db, _ := database.GetDB()

	socialAccount := auth.SocialAccount{
		UserId:   1,
		UUID:     utils.GenerateUUIDString(),
		Provider: "GOOGLE",
		SocialId: utils.GenerateUUIDString(),
	}

	socialAccountService := auth.NewSocialAccountService(db)
	lastInsertId, err := socialAccountService.Create(socialAccount)

	assert.Nil(err, nil)
	assert.NotZero(lastInsertId)
}

func Test_SocialAccountService_Create_Fail(t *testing.T) {
	assert := assert.New(t)

	db, _ := database.GetDB()

	socialAccount := auth.SocialAccount{
		UserId:   1,
		UUID:     utils.GenerateUUIDString(),
		Provider: "",
		SocialId: utils.GenerateUUIDString(),
	}

	socialAccountService := auth.NewSocialAccountService(db)
	lastInsertId, err := socialAccountService.Create(socialAccount)

	assert.Error(err)
	assert.Zero(lastInsertId)
}
