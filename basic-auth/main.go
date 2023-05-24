package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	auth
}

type auth struct {
	Username string
	Password string
}

func main() {
	app := new(application)

	app.auth.Username = os.Getenv("username")
	app.auth.Password = os.Getenv("password")

	if app.auth.Password == "" || app.auth.Username == "" {
		log.Fatal("configs missing")
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/authenticated", app.basicAuth(app.authenticated))
	mux.HandleFunc("/unauthenticated", app.unauthenticated)

	srv := &http.Server{
		Addr:         ":4000",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("starting server on %s", srv.Addr)
	err := srv.ListenAndServeTLS("./localhost.pem", "./localhost-key.pem")
	log.Fatal(err)
}

func (app *application) authenticated(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the protected route")
}

func (app *application) unauthenticated(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the unprotected route")
}

func (app application) basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(app.auth.Username))
			expectedPasswordHash := sha256.Sum256([]byte(app.auth.Password))

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
