package internal

import (
	"fmt"
	"html"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "You've made it to %q, also known as Homepage", html.EscapeString(r.URL.Path))
}
func GetGamesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "return all games")
}
func GetGameHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "returning game with ID %s", id)
}
func CreateGamesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post a new game")
}
