package main

import (
	"encoding/json"
	"net/http"
)

type Watchlist struct {
	ID     string
	Name   string   `json:"name"`
	Alerts []*Alert `json:"alerts"`
}
type Alert struct {
	ID       string
	Ticker   string  `json:"ticker"`
	Operator string  `json:"operator"`
	Price    float32 `json:"price"`
}

type GetWatchlistsRequest struct {
	ID string `json:"ID"`
}

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
	if checkMethod(w, r) == false {
		return
	}

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

func GetWatchlists(w http.ResponseWriter, r *http.Request) {
	if checkMethod(w, r) == false {
		return
	}

	var req GetWatchlistsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	watchlists, err := DB_getWatchlists(req)
	if err != nil {
		http.Error(w, "Failed to fetch watchlists", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(watchlists)
}
