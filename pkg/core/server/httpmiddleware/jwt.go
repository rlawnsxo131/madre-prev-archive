package httpmiddleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/pkg/core/logger"
	"github.com/rlawnsxo131/madre-server-v3/pkg/core/server/httpresponse"
	"github.com/rlawnsxo131/madre-server-v3/pkg/core/token"
	"github.com/rs/zerolog"
)

// When the token already exists,
// if an error occurs when reissuing the token,
// only logging is processed so that other functions can be used.
func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		actk, err := r.Cookie(token.ACCESS_TOKEN)
		tokenManager := token.NewManager()

		if err != nil {
			if err != http.ErrNoCookie {
				httpresponse.NewWriter(w, r).Error(
					errors.Wrap(err, "JWT get Access_token error"),
					httpresponse.NewError(
						http.StatusInternalServerError,
					),
				)
				return
			}
		}

		if actk != nil {
			claims, err := tokenManager.Decode(actk.Value)
			if err != nil {
				_, ok := err.(*jwt.ValidationError)
				if ok {
					rftk, err := r.Cookie(token.REFRESH_TOKEN)
					if err != nil {
						if err != http.ErrNoCookie {
							httpresponse.NewWriter(w, r).Error(
								errors.Wrap(err, "JWT get Refresh_token error"),
								httpresponse.NewError(
									http.StatusInternalServerError,
								),
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
							p := token.NewProfile(
								claims.UserId,
								claims.Username,
								claims.PhotoUrl,
							)
							err = tokenManager.GenerateAndSetCookies(p, w)
							if err != nil {
								logger.DefaultLogger.NewLogEntry().Add(func(e *zerolog.Event) {
									e.Err(err).Str("action", "JWT")
								}).Send()
							} else {
								ctx = token.SetProfileCtx(ctx, p)
							}
						}
					}
				}
			} else {
				p := token.NewProfile(
					claims.UserId,
					claims.Username,
					claims.PhotoUrl,
				)
				ctx = token.SetProfileCtx(ctx, p)
			}
		}

		if actk == nil {
			rftk, err := r.Cookie(token.ACCESS_TOKEN)
			if err != nil {
				if err != http.ErrNoCookie {
					httpresponse.NewWriter(w, r).Error(
						errors.Wrap(err, "JWT get Refresh_token error"),
						httpresponse.NewError(
							http.StatusInternalServerError,
						),
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
					p := token.NewProfile(
						claims.UserId,
						claims.Username,
						claims.PhotoUrl,
					)
					err = tokenManager.GenerateAndSetCookies(p, w)
					if err != nil {
						logger.DefaultLogger.NewLogEntry().Add(func(e *zerolog.Event) {
							e.Err(err).Str("action", "JWT")
						}).Send()
					} else {
						ctx = token.SetProfileCtx(ctx, p)
					}
				}
			}
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
