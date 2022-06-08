package authrepository

import "github.com/rlawnsxo131/madre-server-v3/domain/auth"

type socialAccountEntityMapper struct{}

func (e socialAccountEntityMapper) toEntity(sa *auth.SocialAccount) *auth.SocialAccount {
	return &auth.SocialAccount{
		ID:        sa.ID,
		UserID:    sa.UserID,
		Provider:  sa.Provider,
		SocialID:  sa.SocialID,
		CreatedAt: sa.CreatedAt,
		UpdatedAt: sa.UpdatedAt,
	}
}

func (e socialAccountEntityMapper) toModel(sa *auth.SocialAccount) *auth.SocialAccount {
	return &auth.SocialAccount{
		UserID:   sa.UserID,
		Provider: sa.Provider,
		SocialID: sa.SocialID,
	}
}
