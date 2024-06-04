package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "You've made it to %q, also known as Homepage", html.EscapeString(r.URL.Path))
	})
	mux.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Game List")
	})
	log.Fatal(server.ListenAndServe())
}
