package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/Health", Health)
	mux.HandleFunc("/CreateWatchlist", CreateWatchlist)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
