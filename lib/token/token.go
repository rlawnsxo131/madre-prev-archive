package token

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

type authTokenClaims struct {
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
	GetToken() string
	GenerateToken(params GenerateTokenParams) error
	SetTokenCookie(w http.ResponseWriter)
}

type tokenManager struct {
	token string
}

func NewTokenManager() TokenManager {
	return &tokenManager{}
}

func (tm *tokenManager) GetToken() string {
	return tm.token
}

func (tm *tokenManager) GenerateToken(params GenerateTokenParams) error {
	signKey := []byte("madre base")
	claims := authTokenClaims{
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

	tm.token = ss

	return nil
}

func (tm *tokenManager) SetTokenCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:  "Access_token",
		Value: tm.token,
		Path:  "/",
		// Domain:   ".juntae.kim",
		Expires:  time.Now().AddDate(0, 0, 7),
		Secure:   true,
		HttpOnly: true,
		SameSite: 2,
	})
}
