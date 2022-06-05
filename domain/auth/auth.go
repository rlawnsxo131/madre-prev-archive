package auth

import "net/http"

type Controller interface {
	Get() http.HandlerFunc
	Delete() http.HandlerFunc
}

type GoogleController interface {
	PostGoogleCheck() http.HandlerFunc
	PostGoogleSignIn() http.HandlerFunc
	PostGoogleSignUp() http.HandlerFunc
}
