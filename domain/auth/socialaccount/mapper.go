package socialaccount

type entityMapper struct{}

func (e entityMapper) toEntity(sa *SocialAccount) *SocialAccount {
	return &SocialAccount{
		ID:        sa.ID,
		UserID:    sa.UserID,
		Provider:  sa.Provider,
		SocialId:  sa.SocialId,
		CreatedAt: sa.CreatedAt,
		UpdatedAt: sa.UpdatedAt,
	}
}

func (e entityMapper) toModel(sa *SocialAccount) *SocialAccount {
	return &SocialAccount{
		UserID:   sa.UserID,
		Provider: sa.Provider,
		SocialId: sa.SocialId,
	}
}
