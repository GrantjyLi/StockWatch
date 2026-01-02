package main

import (
	"encoding/json"
	"net/http"
)

type CreateWatchlistRequest struct {
	UserID        string    `json:"userID"`
	WatchlistData Watchlist `json:"watchlistData"`
}
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

type GetWatchlistsRequest_t struct {
	ID string `json:"ID"`
}

type DeleteWatchlistsRequest_t struct {
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
	var req CreateWatchlistRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = DB_writeWatchlist(req.UserID, req.WatchlistData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "watchlist created",
	})
}

func DeleteWatchlists(w http.ResponseWriter, r *http.Request) {
	if checkMethod(w, r) == false {
		return
	}

	var req DeleteWatchlistsRequest_t
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = DB_deleteWatchlist(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "watchlist deleted",
	})
}

func GetWatchlists(w http.ResponseWriter, r *http.Request) {
	if checkMethod(w, r) == false {
		return
	}

	var req GetWatchlistsRequest_t
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
