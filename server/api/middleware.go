package main

import (
	"errors"
	"net/http"

	"boundsApp.victorinolavida/internal/data"
)

func (app *application) enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set the CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) authenticate(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := app.getCookieByName(r, TokenName)

		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				app.invalidCredentialsResponse(w, r)
				return
			}
			app.serverErrorResponse(w, r, err)
			return
		}

		claims := &CustomClaims{}
		err = app.validateToken(cookie, claims)

		if err != nil {

			switch {
			case errors.Is(err, ErrInvalidCredentials):
				app.invalidCredentialsResponse(w, r)
				return
			default:
				app.serverErrorResponse(w, r, err)
				return
			}
		}

		user, err := app.models.Users.GetByUsername(claims.Username)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):
				app.invalidCredentialsResponse(w, r)
				return
			default:
				app.serverErrorResponse(w, r, err)
				return
			}
		}

		r = app.contextSetUser(r, user)
		next.ServeHTTP(w, r)
	}

}