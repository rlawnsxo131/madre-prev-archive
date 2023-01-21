package mapper

import (
	"time"

	"github.com/rlawnsxo131/madre-server-v3/core/utils"
	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
	"github.com/rlawnsxo131/madre-server-v3/internal/infrastructure/persistence/model"
)

type SocialAccountMapper struct{}

func (sam SocialAccountMapper) MapToModel(sa *user.SocialAccount) *model.SocialAccount {
	return &model.SocialAccount{
		Id:             sa.Id,
		UserId:         sa.UserId,
		SocialId:       sa.SocialId,
		SocialUsername: utils.NewNullString(sa.SocialUsername),
		Provider:       sa.Provider,
		UpdatedAt:      time.Now(),
	}
}

func (sam SocialAccountMapper) MapToEntity(sa *model.SocialAccount) *user.SocialAccount {
	return &user.SocialAccount{
		Id:             sa.Id,
		UserId:         sa.UserId,
		SocialId:       sa.SocialId,
		SocialUsername: utils.NormalizeNullString(sa.SocialUsername),
		Provider:       sa.Provider,
	}
}
