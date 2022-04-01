package token

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

type authTokenClaims struct {
	TokenUUID   string `json:"token_uuid"`
	UserID      string `json:"id"`
	UserUUID    string `json:"uuid"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	jwt.StandardClaims
}

type GenerateTokenParams struct {
	UserID      string `json:"id"`
	UserUUID    string `json:"uuid"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type TokenManager interface {
	GetTokens() (string, string)
	GenerateToken(params GenerateTokenParams) error
	SetTokenCookie(w http.ResponseWriter)
}

type tokenManager struct {
	accessToken  string
	refreshToken string
}

func NewTokenManager() TokenManager {
	return &tokenManager{}
}

func (tm *tokenManager) GetTokens() (string, string) {
	return tm.accessToken, tm.refreshToken
}

func (tm *tokenManager) GenerateToken(params GenerateTokenParams) error {
	signKey := []byte("madre base")

	for i := 0; i < 2; i++ {
		claims := authTokenClaims{
			TokenUUID:   utils.GenerateUUIDString(),
			UserID:      params.UserID,
			UserUUID:    params.UserUUID,
			DisplayName: params.DisplayName,
			Email:       params.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: 60 * 60 * 24 * 7,
				Issuer:    "madre",
				IssuedAt:  time.Now().Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(signKey)

		if err != nil {
			return errors.Wrap(err, "GenerateToken")
		}

		if i == 0 {
			tm.accessToken = ss
		} else {
			tm.refreshToken = ss
		}
	}

	return nil
}

func (tm *tokenManager) SetTokenCookie(w http.ResponseWriter) {
	now := time.Now()
	http.SetCookie(w, &http.Cookie{
		Name:  "Refresh_token",
		Value: tm.refreshToken,
		Path:  "/",
		// Domain:   ".juntae.kim",
		Expires:  now.AddDate(0, 0, 30),
		Secure:   true,
		HttpOnly: true,
		SameSite: 2,
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "Access_token",
		Value: tm.accessToken,
		Path:  "/",
		// Domain:   ".juntae.kim",
		Expires:  now.AddDate(0, 0, 7),
		Secure:   true,
		HttpOnly: true,
		SameSite: 2,
	})
}
