package middleware

import "net/http"

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
