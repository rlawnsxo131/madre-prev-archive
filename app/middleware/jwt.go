package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/lib/logger"
	"github.com/rlawnsxo131/madre-server-v3/lib/response"
	"github.com/rlawnsxo131/madre-server-v3/lib/token"
)

// When the token already exists,
// if an error occurs when reissuing the token,
// only logging is processed so that other functions can be used.
func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		actk, err := r.Cookie(token.Key_AccessToken)
		tokenManager := token.NewManager()
		if err != nil {
			if err != http.ErrNoCookie {
				rw := response.NewWriter(w, r)
				rw.Error(
					errors.Wrap(err, "JWT get Access_token error"),
				)
				return
			}
		}

		if actk != nil {
			claims, err := tokenManager.Decode(actk.Value)
			if err != nil {
				_, ok := err.(*jwt.ValidationError)
				if ok {
					rftk, err := r.Cookie(token.Key_RefreshToken)
					if err != nil {
						if err != http.ErrNoCookie {
							rw := response.NewWriter(w, r)
							rw.Error(
								errors.Wrap(err, "JWT get Refresh_token error"),
							)
							return
						}
					}

					if rftk != nil {
						claims, err := tokenManager.Decode(rftk.Value)
						if err != nil {
							// remove cookies
							tokenManager.ResetCookies(w)
						} else {
							p := token.UserProfile{
								UserID:   claims.UserID,
								Username: claims.Username,
								PhotoUrl: claims.PhotoUrl,
							}
							err = tokenManager.GenerateAndSetCookies(&p, w)
							if err != nil {
								logger.DefaultLogger().Err(err).Timestamp().Str("action", "JWT").Send()
							} else {
								ctx = token.SetUserProfileCtx(ctx, &p)
							}
						}
					}
				}
			} else {
				p := token.UserProfile{
					UserID:   claims.UserID,
					Username: claims.Username,
					PhotoUrl: claims.PhotoUrl,
				}
				ctx = token.SetUserProfileCtx(ctx, &p)
			}
		}

		if actk == nil {
			rftk, err := r.Cookie(token.Key_AccessToken)
			if err != nil {
				if err != http.ErrNoCookie {
					rw := response.NewWriter(w, r)
					rw.Error(
						errors.Wrap(err, "get Refresh_token error"),
					)
					return
				}
			}

			if rftk != nil {
				claims, err := tokenManager.Decode(rftk.Value)
				if err != nil {
					// remove cookies
					tokenManager.ResetCookies(w)
				} else {
					p := token.UserProfile{
						UserID:   claims.UserID,
						Username: claims.Username,
						PhotoUrl: claims.PhotoUrl,
					}
					err = tokenManager.GenerateAndSetCookies(&p, w)
					if err != nil {
						logger.DefaultLogger().Err(err).Timestamp().Str("action", "JWT").Send()
					} else {
						ctx = token.SetUserProfileCtx(ctx, &p)
					}
				}
			}
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
