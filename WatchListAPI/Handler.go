package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type CreateWatchlistRequest struct {
	name    string            `json:"name"`
	tickers map[string]string `json:"tickers"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func CreateWatchlist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateWatchlistRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// later: insert into DB
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "watchlist created",
	})

	fmt.Printf("%s:\n", req.name)
	for ticker, condition := range req.tickers {
		fmt.Printf("    %s: %s\n", ticker, condition) 
	}
}
