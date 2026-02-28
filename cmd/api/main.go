package main

import (
	"encoding/json"
	"github/idbeholdv18/go-di/internal/middleware"
	"github/idbeholdv18/go-di/internal/mock"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(&mock.User{
			Name: "Alim",
			Age:  24,
		})
	}))

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatal("stating server error:", err)
	}
}

func NewUnauthorizedError() {

}
