package main

import (
	"database/sql"
	"log"
	"net/http"
)

// API create watchlist test
// curl -Method POST "http://localhost:8080/CreateWatchlist" -Headers @{ "Content-Type" = "application/json" } -Body '{"name":"newListTest","alerts":[{"ticker": "VOO", "condition": "<=", "price": 50}, {"ticker": "VFV", "condition": ">=", "price": 1150}]}'
// API get watchlist test
// curl -Method POST "http://localhost:8080/GetWatchlists" -Headers @{ "Content-Type" = "application/json" } -Body '{"ID": "eb0dcdff-741d-437c-ad64-35b267a91494"}'
var database *sql.DB

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow your frontend origin
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	database = DB_connect()
	defer database.Close()

	mux := http.NewServeMux()

	mux.Handle("/Health", enableCORS(http.HandlerFunc(Health)))
	mux.Handle("/CreateWatchlist", enableCORS(http.HandlerFunc(CreateWatchlist)))
	mux.Handle("/GetWatchlists", enableCORS(http.HandlerFunc(GetWatchlists)))
	mux.Handle("/DeleteWatchlist", enableCORS(http.HandlerFunc(DeleteWatchlists)))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
