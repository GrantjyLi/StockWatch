package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type CreateWatchlistRequest_t struct {
	UserID        string    `json:"userID"`
	User_email    string    `json:"email"`
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

type LoginRequest_t struct {
	User_email string `json:"email"`
}

type CreateUserRequest_t struct {
	User_email string `json:"email"`
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

func LoginRequest(w http.ResponseWriter, r *http.Request) {
	if checkMethod(w, r) == false {
		return
	}

	var req LoginRequest_t
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	userID, err := DB_CheckLogin(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"userID": userID,
	})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if checkMethod(w, r) == false {
		return
	}

	var req CreateUserRequest_t
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = DB_createUser(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "watchlist created",
	})
}

func CreateWatchlist(w http.ResponseWriter, r *http.Request) {
	if checkMethod(w, r) == false {
		return
	}
	var req CreateWatchlistRequest_t
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = DB_createWatchlist(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = publishNewWatchlist(&req)
	if err != nil {
		log.Printf("Error publishing watchlist email: %s\n", err.Error())
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

	err = DB_deleteWatchlist(&req)
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

	watchlists, err := DB_getWatchlists(&req)
	if err != nil {
		http.Error(w, "Failed to fetch watchlists", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(watchlists)
}
