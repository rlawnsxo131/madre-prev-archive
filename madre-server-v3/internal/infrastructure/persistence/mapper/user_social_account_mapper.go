package mapper

import (
	"time"

	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
	"github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/persistence/model"
	valueutil "github.com/rlawnsxo131/madre-server-v3/pkg/core/utils/value-util"
)

type UserSocialAccountMapper struct{}

func (sam UserSocialAccountMapper) MapToModel(sa *user.UserSocialAccount) *model.SocialAccount {
	return &model.SocialAccount{
		Id:             sa.Id,
		UserId:         sa.UserId,
		SocialId:       sa.SocialId,
		SocialUsername: valueutil.NewNullString(sa.SocialUsername),
		Provider:       sa.Provider,
		UpdatedAt:      time.Now(),
	}
}

func (sam UserSocialAccountMapper) MapToEntity(sa *model.SocialAccount) *user.UserSocialAccount {
	return &user.UserSocialAccount{
		Id:             sa.Id,
		UserId:         sa.UserId,
		SocialId:       sa.SocialId,
		SocialUsername: valueutil.NormalizeNullString(sa.SocialUsername),
		Provider:       sa.Provider,
	}
}
