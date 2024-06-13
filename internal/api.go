package internal

import (
	"log"
	"net/http"
)

func InitApi() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("GET /games", GetGamesHandler)
	mux.HandleFunc("GET /game", GetGameHandler)
	mux.HandleFunc("POST /games", CreateGamesHandler)

	log.Fatal(server.ListenAndServe())
}
