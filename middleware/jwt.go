package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/rlawnsxo131/madre-server-v2/lib/httpcontext"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
	"github.com/rlawnsxo131/madre-server-v2/lib/token"
)

// When the token already exists,
// if an error occurs when reissuing the token,
// only logging is processed so that other functions can be used.
func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		atk, err := r.Cookie(token.Key_AccessToken)
		if err != nil {
			if err != http.ErrNoCookie {
				rw := response.NewWriter(w, r)
				rw.Error(
					err,
					"JwtMiddleware",
					"get Access_token error",
				)
				return
			}
		}

		if atk != nil {
			claims, err := token.DecodeToken(atk.Value)
			if err != nil {
				_, ok := err.(*jwt.ValidationError)
				if ok {
					rtk, err := r.Cookie(token.Key_RefreshToken)
					if err != nil {
						if err != http.ErrNoCookie {
							rw := response.NewWriter(w, r)
							rw.Error(
								err,
								"JwtMiddleware",
								"get Refresh_token error",
							)
							return
						}
					}

					if rtk != nil {
						claims, err := token.DecodeToken(rtk.Value)
						if err != nil {
							// remove cookies
							token.ResetTokenCookies(w)

							// set context value
							cm := httpcontext.NewContextManager(ctx)
							ctx = cm.SetUserProfile(nil)
						} else {
							// generate tokens and set cookie
							p := token.UserProfile{
								UserID:      claims.UserID,
								DisplayName: claims.DisplayName,
								PhotoUrl:    claims.PhotoUrl,
							}
							atk, rtk, err := token.GenerateTokens(&p)
							if err != nil {
								logger.NewDefaultLogger().Err(err).Str("Action", "JWT").Msg("")
							} else {
								token.SetTokenCookies(w, atk, rtk)

								// set context value
								cm := httpcontext.NewContextManager(ctx)
								p := token.UserProfile{
									UserID:      claims.UserID,
									DisplayName: claims.DisplayName,
									PhotoUrl:    claims.PhotoUrl,
								}
								ctx = cm.SetUserProfile(&p)
							}
						}
					}
				}
			} else {
				// set context value
				cm := httpcontext.NewContextManager(ctx)
				p := token.UserProfile{
					UserID:      claims.UserID,
					DisplayName: claims.DisplayName,
					PhotoUrl:    claims.PhotoUrl,
				}
				ctx = cm.SetUserProfile(&p)
			}
		}

		if atk == nil {
			rtk, err := r.Cookie(token.Key_AccessToken)
			if err != nil {
				if err != http.ErrNoCookie {
					rw := response.NewWriter(w, r)
					rw.Error(
						err,
						"JwtMiddleware",
						"get Refresh_token error",
					)
					return
				}
			}

			if rtk != nil {
				claims, err := token.DecodeToken(rtk.Value)
				if err != nil {
					// remove cookies
					token.ResetTokenCookies(w)
					cm := httpcontext.NewContextManager(ctx)
					ctx = cm.SetUserProfile(nil)

				} else {
					// generate tokens and set cookie
					p := token.UserProfile{
						UserID:      claims.UserID,
						DisplayName: claims.DisplayName,
						PhotoUrl:    claims.PhotoUrl,
					}
					atk, rtk, err := token.GenerateTokens(&p)
					if err != nil {
						logger.NewDefaultLogger().Err(err).Str("Action", "JWT").Msg("")
					} else {
						token.SetTokenCookies(w, atk, rtk)
						cm := httpcontext.NewContextManager(ctx)
						ctx = cm.SetUserProfile(nil)
					}
				}
			}
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
