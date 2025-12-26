package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

const ENV_FILE = "WatchListAPI.env"

func main() {

	err := godotenv.Load(ENV_FILE)
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	DB_USER := os.Getenv("DATABASE_USER")
	DB_PW := os.Getenv("DATABASE_PASSWORD")
	DB_NAME := os.Getenv("DATABASE_NAME")
	DB_ENDPOINT := os.Getenv("DATABASE_ENDPOINT")
	DB_PORT := os.Getenv("DATABASE_PORT")

	DB_DSN := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=require", DB_USER, DB_PW, DB_ENDPOINT, DB_PORT, DB_NAME)

	db, err := sql.Open("pgx", DB_DSN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open database connection: %v\n", err)
		os.Exit(1)
	}

	// db.Close() is deferred to ensure the connection is closed when the main function exits.
	defer db.Close()

	// Call db.Ping() to verify that the database connection is alive and establishable.
	pingErr := db.Ping()
	if pingErr != nil {
		// log.Fatal prints the error and then calls os.Exit(1).
		log.Fatal(pingErr)
	}

	fmt.Println("Successfully connected to the database!")
}