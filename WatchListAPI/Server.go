package main

import (
	"database/sql"
	"log"
	"net/http"
)

var database *sql.DB
func main() {
	database = DB_connect()
	defer database.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/Health", Health)
	mux.HandleFunc("/CreateWatchlist", CreateWatchlist)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
