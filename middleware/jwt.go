package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/rlawnsxo131/madre-server-v2/constants"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
	"github.com/rlawnsxo131/madre-server-v2/lib/syncmap"
	"github.com/rlawnsxo131/madre-server-v2/lib/token"
)

// When the token already exists,
// if an error occurs when reissuing the token,
// only logging is processed so that other functions can be used.
func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := r.Cookie(token.Key_AccessToken)
		if err != nil {
			if err != http.ErrNoCookie {
				writer := response.NewHttpWriter(w, r)
				writer.Error(
					err,
					"JwtMiddleware",
					"get Access_token error",
				)
				return
			}
		}

		if accessToken != nil {
			claims, err := token.DecodeToken(accessToken.Value)
			if err != nil {
				_, ok := err.(*jwt.ValidationError)
				if ok {
					refreshToken, err := r.Cookie(token.Key_RefreshToken)
					if err != nil {
						if err != http.ErrNoCookie {
							writer := response.NewHttpWriter(w, r)
							writer.Error(
								err,
								"JwtMiddleware",
								"get Refresh_token error",
							)
							return
						}
					}

					if refreshToken != nil {
						claims, err := token.DecodeToken(refreshToken.Value)
						if err != nil {
							// remove cookies
							token.ResetTokenCookies(w)

							// set context value
							ctx, err := syncmap.SetNewValueFromHttpContext(
								r.Context(),
								constants.Key_UserTokenProfile,
								nil,
							)
							if err != nil {
								writer := response.NewHttpWriter(w, r)
								writer.Error(
									err,
									"JWT",
								)
								return
							}
							r.Context().Value(ctx)
						} else {
							// generate tokens and set cookie
							profile := token.UserTokenProfile{
								UserID:      claims.UserID,
								DisplayName: claims.DisplayName,
								PhotoUrl:    claims.PhotoUrl,
							}
							accessToken, refreshToken, err := token.GenerateTokens(&profile)
							if err != nil {
								logger.NewDefaultLogger().
									Err(err).
									Str("Action", "JWT").
									Send()
							} else {
								token.SetTokenCookies(w, accessToken, refreshToken)

								// set context value
								ctx, err := syncmap.SetNewValueFromHttpContext(
									r.Context(),
									constants.Key_UserTokenProfile,
									&token.UserTokenProfile{
										UserID:      claims.UserID,
										DisplayName: claims.DisplayName,
										PhotoUrl:    claims.PhotoUrl,
									},
								)
								if err != nil {
									writer := response.NewHttpWriter(w, r)
									writer.Error(
										err,
										"JWT",
									)
									return
								}
								r.Context().Value(ctx)
							}
						}
					}
				}
			} else {
				// set context value
				ctx, err := syncmap.SetNewValueFromHttpContext(
					r.Context(),
					constants.Key_UserTokenProfile,
					&token.UserTokenProfile{
						UserID:      claims.UserID,
						DisplayName: claims.DisplayName,
						PhotoUrl:    claims.PhotoUrl,
					},
				)
				if err != nil {
					writer := response.NewHttpWriter(w, r)
					writer.Error(
						err,
						"JWT",
					)
					return
				}
				r.Context().Value(ctx)
			}
		}

		if accessToken == nil {
			refreshToken, err := r.Cookie(token.Key_RefreshToken)
			if err != nil {
				if err != http.ErrNoCookie {
					writer := response.NewHttpWriter(w, r)
					writer.Error(
						err,
						"JwtMiddleware",
						"get Refresh_token error",
					)
					return
				}
			}

			if refreshToken != nil {
				claims, err := token.DecodeToken(refreshToken.Value)
				if err != nil {
					// remove cookies
					token.ResetTokenCookies(w)

					// set context value
					ctx, err := syncmap.SetNewValueFromHttpContext(
						r.Context(),
						constants.Key_UserTokenProfile,
						nil,
					)
					if err != nil {
						writer := response.NewHttpWriter(w, r)
						writer.Error(
							err,
							"JWT",
						)
						return
					}
					r.Context().Value(ctx)
				} else {
					// generate tokens and set cookie
					profile := token.UserTokenProfile{
						UserID:      claims.UserID,
						DisplayName: claims.DisplayName,
						PhotoUrl:    claims.PhotoUrl,
					}
					accessToken, refreshToken, err := token.GenerateTokens(&profile)
					if err != nil {
						logger.NewDefaultLogger().
							Err(err).
							Str("Action", "JWT").
							Send()
					} else {
						token.SetTokenCookies(w, accessToken, refreshToken)

						// set context value
						ctx, err := syncmap.SetNewValueFromHttpContext(
							r.Context(),
							constants.Key_UserTokenProfile,
							&token.UserTokenProfile{
								UserID:      claims.UserID,
								DisplayName: claims.DisplayName,
								PhotoUrl:    claims.PhotoUrl,
							},
						)
						if err != nil {
							writer := response.NewHttpWriter(w, r)
							writer.Error(
								err,
								"JWT",
							)
							return
						}
						r.Context().Value(ctx)
					}
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}
