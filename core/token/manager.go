package token

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/core/env"
)

const (
	ACCESS_TOKEN  = "Access_token"
	REFRESH_TOKEN = "Refresh_token"
)

var (
	tokenTypes = []string{ACCESS_TOKEN, REFRESH_TOKEN}
)

type authTokenClaims struct {
	TokenUUID string `json:"token_uuid"`
	UserId    string `json:"user_id"`
	Username  string `json:"username"`
	PhotoUrl  string `json:"photo_url"`
	jwt.RegisteredClaims
}

type manager struct{}

func NewManager() *manager {
	return &manager{}
}

func (m *manager) GenerateAndSetCookies(p *profile, w http.ResponseWriter) error {
	actk, rftk, err := m.generateTokens(p)
	if err != nil {
		return err
	}
	m.setCookies(w, actk, rftk)
	return nil
}

func (m *manager) Decode(token string) (*authTokenClaims, error) {
	claims := authTokenClaims{}
	t, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(env.JWTSecretKey()), nil
		}
		return nil, errors.New("Docode: ParseWithClaims")
	})

	if err != nil {
		return nil, err
	}

	if t.Valid {
		return &claims, nil
	}

	return nil, errors.New("Decode: token is not valid")
}

func (m *manager) ResetCookies(w http.ResponseWriter) {
	now := time.Now()
	expires := now.AddDate(0, 0, -1)

	http.SetCookie(w, &http.Cookie{
		Name:  ACCESS_TOKEN,
		Value: "",
		Path:  "/",
		// Domain:   ".juntae.kim",
		Expires:  expires,
		Secure:   true,
		HttpOnly: true,
		SameSite: 2,
	})
	http.SetCookie(w, &http.Cookie{
		Name:  REFRESH_TOKEN,
		Value: "",
		Path:  "/",
		// Domain:   ".juntae.kim",
		Expires:  expires,
		Secure:   true,
		HttpOnly: true,
		SameSite: 2,
	})
}

func (m *manager) generateTokens(p *profile) (string, string, error) {
	var actk string
	var rftk string

	now := time.Now()
	claimsFuncMap := map[string]jwt.RegisteredClaims{
		ACCESS_TOKEN: {
			ExpiresAt: jwt.NewNumericDate(now.AddDate(0, 0, 1)),
			Issuer:    "madre",
			IssuedAt:  jwt.NewNumericDate(now),
		},
		REFRESH_TOKEN: {
			ExpiresAt: jwt.NewNumericDate(now.AddDate(0, 0, 7)),
			Issuer:    "madre",
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	claims := &authTokenClaims{
		UserId:   p.UserId,
		Username: p.Username,
		PhotoUrl: p.PhotoUrl,
	}

	for _, tokenType := range tokenTypes {
		claims.TokenUUID = uuid.NewString()
		claims.RegisteredClaims = claimsFuncMap[tokenType]

		t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := t.SignedString([]byte(env.JWTSecretKey()))
		if err != nil {
			return "", "", errors.Wrap(err, "generateTokens")
		}

		if tokenType == ACCESS_TOKEN {
			actk = ss
			continue
		}
		rftk = ss
	}

	return actk, rftk, nil
}

func (m *manager) setCookies(w http.ResponseWriter, actk, rftk string) {
	now := time.Now()

	http.SetCookie(w, &http.Cookie{
		Name:  ACCESS_TOKEN,
		Value: actk,
		Path:  "/",
		// Domain:   ".juntae.kim",
		Expires:  now.AddDate(0, 0, 1),
		Secure:   true,
		HttpOnly: true,
		SameSite: 2,
	})
	http.SetCookie(w, &http.Cookie{
		Name:  REFRESH_TOKEN,
		Value: rftk,
		Path:  "/",
		// Domain:   ".juntae.kim",
		Expires:  now.AddDate(0, 0, 7),
		Secure:   true,
		HttpOnly: true,
		SameSite: 2,
	})
}
