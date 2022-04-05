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

const (
	AccessToken  = "Access_token"
	RefreshToken = "Refresh_token"
)

var (
	signKey    = []byte("madre base")
	tokenTypes = []string{AccessToken, RefreshToken}
)

func GenerateAccessToken(params GenerateTokenParams) (string, error) {
	now := time.Now()

	claims := authTokenClaims{
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signKey)
	if err != nil {
		return "", errors.Wrap(err, "GenerateToken")
	}

	return ss, nil
}

func GenerateTokens(params GenerateTokenParams) (string, string, error) {
	now := time.Now()
	var accessToken string
	var refreshToken string

	for _, tokenType := range tokenTypes {
		var claims authTokenClaims

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
				TokenUUID:   utils.GenerateUUIDString(),
				UserUUID:    params.UserUUID,
				DisplayName: params.DisplayName,
				Email:       params.Email,
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
			return "", "", errors.Wrap(err, "GenerateToken")
		}

		if tokenType == AccessToken {
			accessToken = ss
			continue
		}
		refreshToken = ss
	}

	return accessToken, refreshToken, nil
}

func DecodeToken(token string) (*authTokenClaims, error) {
	claims := authTokenClaims{}
	t, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return signKey, nil
		}
		return nil, errors.New("ParseWithClaims error")
	})

	if err != nil {
		return nil, err
	}

	if t.Valid {
		return &claims, nil
	}

	return nil, errors.New("DecodeToken: token is not valid")
}

func SetTokenCookieAccessToken(w http.ResponseWriter, accessToken string) {
	now := time.Now()

	http.SetCookie(w, &http.Cookie{
		Name:  AccessToken,
		Value: accessToken,
		Path:  "/",
		// Domain:   ".juntae.kim",
		Expires:  now.AddDate(0, 0, 7),
		Secure:   true,
		HttpOnly: true,
		SameSite: 2,
	})
}

func SetTokenCookies(w http.ResponseWriter, accessToken string, refreshToken string) {
	now := time.Now()

	http.SetCookie(w, &http.Cookie{
		Name:  AccessToken,
		Value: accessToken,
		Path:  "/",
		// Domain:   ".juntae.kim",
		Expires:  now.AddDate(0, 0, 7),
		Secure:   true,
		HttpOnly: true,
		SameSite: 2,
	})
	http.SetCookie(w, &http.Cookie{
		Name:  RefreshToken,
		Value: refreshToken,
		Path:  "/",
		// Domain:   ".juntae.kim",
		Expires:  now.AddDate(0, 0, 30),
		Secure:   true,
		HttpOnly: true,
		SameSite: 2,
	})
}
