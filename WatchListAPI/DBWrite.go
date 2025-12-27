package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

const ENV_FILE = "WatchListAPI.env"

func DB_connect() *sql.DB {
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

	// ping to verify that the database connection is alive/establishable
	pingErr := db.Ping()
	if pingErr != nil {
		// log.Fatal prints the error and then calls os.Exit(1).
		log.Fatal(pingErr)
	}

	return db
}

func DB_writeWatchlist(watchlistData CreateWatchlistRequest) {
	userID := "eb0dcdff-741d-437c-ad64-35b267a91494"
	watchlistID := uuid.New()

	_, err = tx.Exec(
		`INSERT INTO watchlists (id, user_id, name)
		VALUES ($1, $2, $3)`,
		watchlistID,
		userID,
		watchlistData.Name,
	)
	if err != nil {
		log.Fatal(err)
	}


	alertID := uuid.New()
	_, err = tx.Exec(
		`INSERT INTO alerts (id, watchlist_id, ticker, operator, target_price)
		VALUES ($1, $2, $3, $4, $5)`,
		alertID,
		watchlistID
		userID,
		watchlistName,
	)
	if err != nil {
		log.Fatal(err)
	}
}