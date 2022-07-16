package mapper

import (
	"database/sql"
	"time"

	"github.com/rlawnsxo131/madre-server-v3/internal/domain/user"
	"github.com/rlawnsxo131/madre-server-v3/utils"
)

type SocialAccountModel struct {
	Id             string         `db:"id"`
	UserId         string         `db:"user_id"`
	SocialId       string         `db:"social_id"`
	SocialUsername sql.NullString `db:"social_username"`
	Provider       string         `db:"provider"`
	CreatedAt      time.Time      `db:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at"`
}

type SocialAccountMapper struct{}

func (sam SocialAccountMapper) MapToModel(sa *user.SocialAccount) *SocialAccountModel {
	return &SocialAccountModel{
		Id:             sa.Id,
		UserId:         sa.UserId,
		SocialId:       sa.SocialId,
		SocialUsername: utils.NewNullString(sa.SocialUsername),
		Provider:       sa.Provider,
		UpdatedAt:      time.Now(),
	}
}

func (sam SocialAccountMapper) MapToEntity(sa *SocialAccountModel) *user.SocialAccount {
	return &user.SocialAccount{
		Id:             sa.Id,
		UserId:         sa.UserId,
		SocialId:       sa.SocialId,
		SocialUsername: utils.NormalizeNullString(sa.SocialUsername),
		Provider:       sa.Provider,
	}
}
