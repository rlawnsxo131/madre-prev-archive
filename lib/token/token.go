package token

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/constants"
	"github.com/rlawnsxo131/madre-server-v2/utils"
)

const (
	Key_AccessToken  = "Access_token"
	Key_RefreshToken = "Refresh_token"
)

var (
	signKey    = []byte("madre base")
	tokenTypes = []string{Key_AccessToken, Key_RefreshToken}
)

type UserTokenProfile struct {
	DisplayName string `json:"display_name"`
	PhotoUrl    string `json:"photo_url"`
	AccessToken string `json:"access_token"`
}

type authTokenClaims struct {
	TokenUUID   string `json:"token_uuid"`
	UserID      string `json:"user_id"`
	DisplayName string `json:"display_name"`
	PhotoUrl    string `json:"photo_url"`
	jwt.StandardClaims
}

func GenerateTokens(userId string, displayName string, photoUrl string) (string, string, error) {
	now := time.Now()
	var accessToken string
	var refreshToken string

	for _, tokenType := range tokenTypes {
		var claims authTokenClaims

		if tokenType == Key_AccessToken {
			claims = authTokenClaims{
				TokenUUID:   utils.GenerateUUIDString(),
				UserID:      userId,
				DisplayName: displayName,
				PhotoUrl:    photoUrl,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: now.AddDate(0, 0, 1).Unix(),
					Issuer:    "madre",
					IssuedAt:  now.Unix(),
				},
			}
		}
		if tokenType == Key_RefreshToken {
			claims = authTokenClaims{
				TokenUUID:   utils.GenerateUUIDString(),
				UserID:      userId,
				DisplayName: displayName,
				PhotoUrl:    photoUrl,
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

		if tokenType == Key_AccessToken {
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

func SetTokenCookies(w http.ResponseWriter, accessToken string, refreshToken string) {
	now := time.Now()

	http.SetCookie(w, &http.Cookie{
		Name:  Key_AccessToken,
		Value: accessToken,
		Path:  "/",
		// Domain:   ".juntae.kim",
		Expires:  now.AddDate(0, 0, 1),
		Secure:   true,
		HttpOnly: true,
		SameSite: 2,
	})
	http.SetCookie(w, &http.Cookie{
		Name:  Key_RefreshToken,
		Value: refreshToken,
		Path:  "/",
		// Domain:   ".juntae.kim",
		Expires:  now.AddDate(0, 0, 30),
		Secure:   true,
		HttpOnly: true,
		SameSite: 2,
	})
}

func ResetTokenCookies(w http.ResponseWriter) {
	now := time.Now()
	expires := now.AddDate(0, 0, -1)

	http.SetCookie(w, &http.Cookie{
		Name:  Key_AccessToken,
		Value: "",
		Path:  "/",
		// Domain:   ".juntae.kim",
		Expires:  expires,
		Secure:   true,
		HttpOnly: true,
		SameSite: 2,
	})
	http.SetCookie(w, &http.Cookie{
		Name:  Key_RefreshToken,
		Value: "",
		Path:  "/",
		// Domain:   ".juntae.kim",
		Expires:  expires,
		Secure:   true,
		HttpOnly: true,
		SameSite: 2,
	})
}

func LoadUserTokenProfileFromHttpContextSyncMap(ctx context.Context) *UserTokenProfile {
	v := ctx.Value(constants.Key_HttpSyncMap)
	syncMap, ok := v.(*sync.Map)

	if ok {
		if profile, ok := syncMap.Load(constants.Key_UserTokenProfile); ok {
			if profile, ok := profile.(*UserTokenProfile); ok {
				return profile
			} else {
				return nil
			}
		} else {
			return nil
		}
	}

	return nil
}
