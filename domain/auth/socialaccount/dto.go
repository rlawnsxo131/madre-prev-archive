package socialaccount

type (
	SocialIDAndProviderDto struct {
		SocialID string `db:"social_id" validate:"required"`
		Provider string `db:"provider" validate:"required"`
	}
)
