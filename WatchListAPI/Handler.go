package main

import (
	"encoding/json"
	"net/http"
)

type Watchlist struct {
	Name    string            `json:"name"`
	Tickers map[string]string `json:"tickers"`
}
/*
{
	"name": "watchlist-name",
	"tickers": {
		"AAPL":">= 150",
		"MSFT":"<= 300"
	}
}
*/
type GetWatchlistsRequest struct {
	ID    string            `json:"ID"`
}
/*
{
	"ID": "user-UUID",
}
*/

func checkMethod(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func CreateWatchlist(w http.ResponseWriter, r *http.Request) {
	if checkMethod(w, r) == false {return}

	var req Watchlist
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	DB_writeWatchlist(req)
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "watchlist created",
	})
}

// func GetWatchlists(w http.ResponseWriter, r *http.Request) {
//     if checkMethod(w, r) == false {return}

// 	var req GetWatchlistsRequest
// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 		return
// 	}

//     // Fetch from DB (you implement this)
//     DB_getWatchlists(req)

//     w.Header().Set("Content-Type", "application/json")
//     json.NewEncoder(w).Encode(watchlists)
// }

