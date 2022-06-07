package repository

import "github.com/rlawnsxo131/madre-server-v2/domain_v2/auth"

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
