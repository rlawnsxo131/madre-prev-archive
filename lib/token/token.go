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
	UserUUID    string `json:"uuid"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type TokenManager interface {
	GetTokens() (string, string)
	GenerateToken(params GenerateTokenParams) error
	DecodeToken(token string) (*authTokenClaims, error)
	SetTokenCookie(w http.ResponseWriter)
}

type tokenManager struct {
	accessToken  string
	refreshToken string
}

const (
	AccessToken  = "Access_token"
	RefreshToken = "Refresh_token"
)

var (
	signKey    = []byte("madre base")
	tokenTypes = []string{AccessToken, RefreshToken}
)

func NewTokenManager() TokenManager {
	return &tokenManager{}
}

func (tm *tokenManager) GetTokens() (string, string) {
	return tm.accessToken, tm.refreshToken
}

func (tm *tokenManager) GenerateToken(params GenerateTokenParams) error {
	for _, tokenType := range tokenTypes {
		var claims authTokenClaims
		now := time.Now()

		if tokenType == AccessToken {
			claims = authTokenClaims{
				TokenUUID:   utils.GenerateUUIDString(),
				UserUUID:    params.UserUUID,
				DisplayName: params.DisplayName,
				Email:       params.Email,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: now.Add(time.Hour * 24).Unix(),
					Issuer:    "madre",
					IssuedAt:  now.Unix(),
				},
			}
		}
		if tokenType == RefreshToken {
			claims = authTokenClaims{
				TokenUUID: utils.GenerateUUIDString(),
				UserUUID:  params.UserUUID,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: now.AddDate(0, 0, 30).Unix(),
					Issuer:    "madre",
					IssuedAt:  now.Unix(),
				},
			}
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(signKey)

		if err != nil {
			return errors.Wrap(err, "GenerateToken")
		}

		if tokenType == AccessToken {
			tm.accessToken = ss
			continue
		}
		tm.refreshToken = ss
	}

	return nil
}

func (tm *tokenManager) DecodeToken(token string) (*authTokenClaims, error) {
	claims := authTokenClaims{}
	t, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return signKey, nil
		}
		return nil, errors.New("ParseWithClaims error")
	})

	if err != nil {
		return nil, errors.Wrap(err, "DecodeToken")
	}

	if t.Valid {
		return &claims, nil
	}

	return nil, errors.New("DecodeToken: token is not valid")
}

func (tm *tokenManager) SetTokenCookie(w http.ResponseWriter) {
	now := time.Now()
	http.SetCookie(w, &http.Cookie{
		Name:  AccessToken,
		Value: tm.accessToken,
		Path:  "/",
		// Domain:   ".juntae.kim",
		Expires:  now.AddDate(0, 0, 7),
		Secure:   true,
		HttpOnly: true,
		SameSite: 2,
	})
	http.SetCookie(w, &http.Cookie{
		Name:  RefreshToken,
		Value: tm.refreshToken,
		Path:  "/",
		// Domain:   ".juntae.kim",
		Expires:  now.AddDate(0, 0, 30),
		Secure:   true,
		HttpOnly: true,
		SameSite: 2,
	})
}
