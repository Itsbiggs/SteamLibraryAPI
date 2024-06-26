package internal

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

type Game struct {
	Title    string `json:"title"`
	ID       string `json:"id"`
	Played   bool   `json:"played"`
	Owned    bool   `json:"owned"`
	Wishlist bool   `json:"wishlist"`
	Hours    int    `json:"hours"`
}

var games = make(map[string]Game)

// handle request to home page, this eventually will do something besides print a statement, TODO
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "You've made it to %q, also known as Homepage", html.EscapeString(r.URL.Path))
}
func GetGamesHandler(w http.ResponseWriter, r *http.Request) {
	//since we are saving to memory we're just going to pull the list, we will need to pull from a DB when implemented.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

// GetGameHandler will let us pull a game based on ID in curl request
func GetGameHandler(w http.ResponseWriter, r *http.Request) {
	//extract ID from query parameters
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required to get game", http.StatusBadRequest)
		return
	}
	if game, ok := games[id]; ok {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(game)
		return
	}
	http.Error(w, "Game not found, are you sure you added it?", http.StatusNotFound)
}

// CreateGamesHandler is a simple post handler that sets the variable game to Game json struct
func CreateGamesHandler(w http.ResponseWriter, r *http.Request) {
	var game Game
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if game.Owned {
		game.Wishlist = false
	}
	if game.Wishlist {
		game.Played = false
		game.Hours = 0
	}
	//add the game to the map by ID
	games[game.ID] = game
	//report success and encode data to json for consumption
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}

func UpdateGameHandler(w http.ResponseWriter, r *http.Request) {
	//catch error of no ID given
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required to update", http.StatusBadRequest)
		return
	}
	//catch error of a bad request
	var game Game
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if game.Owned {
		game.Wishlist = false
	}
	games[id] = game
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "Game updated successfully"}
	json.NewEncoder(w).Encode(response)
	json.NewEncoder(w).Encode(game)
}

func DeleteGameHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required to delete", http.StatusBadRequest)
		return
	}
	var game Game
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	delete(games, id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "Game successfully deleted"}
	json.NewEncoder(w).Encode(response)
}
