package main

import (
	"encoding/json"
	"github/idbeholdv18/go-di/internal/middleware"
	"github/idbeholdv18/go-di/internal/mock"
	"github/idbeholdv18/go-di/internal/service"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	authService := &service.BasicAuthService{}

	authMiddleware := middleware.AuthMiddleware(authService)

	mux.HandleFunc("/hello", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(&mock.User{
			Name: "Alim",
			Age:  24,
		})
	}))

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatal("stating server error:", err)
	}
}
