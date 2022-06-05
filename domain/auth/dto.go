package auth

type (
	PostGoogleCheckRequestDto struct {
		AccessToken string `json:"access_token" validate:"required,min=50"`
	}
	PostGoogleSignInRequestDto struct {
		AccessToken string `json:"access_token" validate:"required,min=50"`
	}
	PostGoogleSignUpRequestDto struct {
		AccessToken string `json:"access_token" validate:"required,min=50"`
		Username    string `json:"username" validate:"required,max=20,min=1"`
	}
)
