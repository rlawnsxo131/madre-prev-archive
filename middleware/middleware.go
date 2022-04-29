package middleware

import (
	"errors"
	"runtime/debug"
	"time"

	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rlawnsxo131/madre-server-v2/constants"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
	"github.com/rlawnsxo131/madre-server-v2/lib/syncmap"
	"github.com/rlawnsxo131/madre-server-v2/lib/token"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				writer := response.NewHttpWriter(w, r)
				writer.WriteError(
					errors.New(string(debug.Stack())),
					"Recovery",
				)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func HttpLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		httpLogger := logger.NewHttpLogger()

		bodyBuffer, reader, err := httpLogger.ReadBody(r.Body)
		if err != nil {
			writer := response.NewHttpWriter(w, r)
			writer.WriteError(
				err,
				"HttpLogger",
			)
			return
		}
		r.Body = reader

		defer func() {
			httpLogger.LogEntry(r, start, string(bodyBuffer))
		}()
		next.ServeHTTP(w, r)
	})
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowHosts := []string{"http://localhost:8080", "http://localhost:5000"}
		origin := r.Header.Get("Origin")
		validation := false
		for _, host := range allowHosts {
			if origin == host {
				validation = true
				break
			}
		}
		if validation {
			for _, method := range []string{"OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"} {
				if method == r.Method {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					w.Header().Set("Access-Control-Allow-Credentials", "true")
					if method == "OPTIONS" {
						w.Header().Set(
							"Access-Control-Allow-Headers",
							"Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, Cookie, X-CSRF-Token",
						)
						w.Header().Set(
							"Access-Control-Allow-Methods",
							"GET,POST,PUT,PATCH,DELETE,OPTIONS,HEAD",
						)
						return
					}
					break
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}

func SetSyncMapContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := syncmap.GenerateHttpContext(r.Context())
		r.Context().Value(ctx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func SetDBContext(db *sqlx.DB) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, err := syncmap.SetNewValueFromHttpContext(
				r.Context(),
				constants.Key_HttpContextDB,
				db,
			)
			if err != nil {
				writer := response.NewHttpWriter(w, r)
				writer.WriteError(
					err,
					"SetDBContext",
					"context set error",
				)
				return
			}

			r.Context().Value(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// When the token already exists,
// if an error occurs when reissuing the token,
// only logging is processed so that other functions can be used.
func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := r.Cookie(token.Key_AccessToken)
		if err != nil {
			if err != http.ErrNoCookie {
				writer := response.NewHttpWriter(w, r)
				writer.WriteError(
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
							writer.WriteError(
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
								writer.WriteError(
									err,
									"JWT",
								)
								return
							}
							r.Context().Value(ctx)
						} else {
							// generate tokens and set cookie
							accessToken, refreshToken, err := token.GenerateTokens(
								claims.UserID,
								claims.DisplayName,
								claims.PhotoUrl,
							)
							if err != nil {
								logger.Logger.
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
										DisplayName: claims.DisplayName,
										PhotoUrl:    claims.PhotoUrl,
										AccessToken: accessToken,
									},
								)
								if err != nil {
									writer := response.NewHttpWriter(w, r)
									writer.WriteError(
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
						DisplayName: claims.DisplayName,
						PhotoUrl:    claims.PhotoUrl,
						AccessToken: accessToken.Value,
					},
				)
				if err != nil {
					writer := response.NewHttpWriter(w, r)
					writer.WriteError(
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
					writer.WriteError(
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
						writer.WriteError(
							err,
							"JWT",
						)
						return
					}
					r.Context().Value(ctx)
				} else {
					// generate tokens and set cookie
					accessToken, refreshToken, err := token.GenerateTokens(
						claims.UserID,
						claims.DisplayName,
						claims.PhotoUrl,
					)
					if err != nil {
						logger.Logger.
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
								DisplayName: claims.DisplayName,
								PhotoUrl:    claims.PhotoUrl,
								AccessToken: accessToken,
							},
						)
						if err != nil {
							writer := response.NewHttpWriter(w, r)
							writer.WriteError(
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

func SetResponseContentTypeJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
