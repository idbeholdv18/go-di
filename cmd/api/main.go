package main

import (
	"encoding/json"
	"github/idbeholdv18/go-di/internal/domain"
	"github/idbeholdv18/go-di/internal/mock"
	"log"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			domain.SendError(w, r, domain.NewUnauthorizedError)
			return
		}

		token := strings.TrimPrefix(authHeader, "Basic")
		if token == "" {
			domain.SendError(w, r, domain.NewUnauthorizedError)
			return
		}

		json.NewEncoder(w).Encode(&mock.User{
			Name: "Alim",
			Age:  24,
		})
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatal("stating server error:", err)
	}
}

func NewUnauthorizedError() {

}
