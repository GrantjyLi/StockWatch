package main

import (
	"database/sql"
	"log"
	"net/http"
)

// API create watchlist test
//curl -Method POST "http://localhost:8080/CreateWatchlist" -Headers @{ "Content-Type" = "application/json" } -Body '{"name":"Tech Stocks","tickers":{"AAPL":">= 150","MSFT":"<= 300"}}'

var database *sql.DB
func main() {
	database = DB_connect()
	defer database.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/Health", Health)
	mux.HandleFunc("/CreateWatchlist", CreateWatchlist)
	mux.HandleFunc("/GetWatchlists", GetWatchlists)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
