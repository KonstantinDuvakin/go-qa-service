package server

import (
	"net/http"

	"github.com/KonstantinDuvakin/go-qa-service/internal/server/handlers"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/version", handlers.Version)

	return mux
}
