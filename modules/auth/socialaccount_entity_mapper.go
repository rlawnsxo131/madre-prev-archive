package auth

type socialAccountEntityMapper struct{}

func (e socialAccountEntityMapper) toEntity(sa *SocialAccount) *SocialAccount {
	return &SocialAccount{
		ID:        sa.ID,
		UserID:    sa.UserID,
		Provider:  sa.Provider,
		SocialID:  sa.SocialID,
		CreatedAt: sa.CreatedAt,
		UpdatedAt: sa.UpdatedAt,
	}
}

func (e socialAccountEntityMapper) toModel(sa *SocialAccount) *SocialAccount {
	return &SocialAccount{
		UserID:   sa.UserID,
		Provider: sa.Provider,
		SocialID: sa.SocialID,
	}
}
