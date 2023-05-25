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

const (
	apiKeyHeader = "X-API-Key"
)

type application struct {
	auth
}

type auth struct {
	APIKey string
}

func main() {
	app := new(application)

	app.auth.APIKey = os.Getenv("apiKey")

	if app.auth.APIKey == "" {
		log.Fatal("configs missing")
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/authenticated", app.apiKeyAuth(app.authenticated))
	mux.HandleFunc("/unauthenticated", app.unauthenticated)

	srv := &http.Server{
		Addr:         ":4001",
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

func (app application) apiKeyAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get(apiKeyHeader)
		if key != "" {
			apiKeyHash := sha256.Sum256([]byte(key))
			expectedapiKeyHash := sha256.Sum256([]byte(app.auth.APIKey))

			apiKeyMatch := (subtle.ConstantTimeCompare(apiKeyHash[:], expectedapiKeyHash[:]) == 1)

			if apiKeyMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
