package main

import (
	"log"
	"net/http"

	"github.com/KonstantinDuvakin/go-qa-service/internal/config"
	"github.com/KonstantinDuvakin/go-qa-service/internal/server"
)

func main() {
	cfg := config.Load()

	router := server.NewRouter()

	port := ":" + cfg.AppPort

	log.Printf("server started on %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
