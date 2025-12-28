package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

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

func DB_writeWatchlist(watchlistData WatchlistRequest) {
	tx, err := database.Begin()
	if err != nil { log.Fatal(err) }
	defer tx.Rollback() // safe rollback if commit never happens

	userID := "eb0dcdff-741d-437c-ad64-35b267a91494"
	watchlistID := uuid.New().String()

	watchlistQuery := fmt.Sprintf("INSERT INTO watchlists (id, user_id, name)\nVALUES ('%s', '%s', '%s');",
		watchlistID,
		userID,
		watchlistData.Name,
	)
	fmt.Println(watchlistQuery)

	_, err = tx.Exec(watchlistQuery)
	if err != nil {
		log.Fatal(err)
	}

	alertsQuery := "INSERT INTO alerts (id, watchlist_id, ticker, operator, target_price)\nValues"
	rows := []string{}
	for ticker, alert := range watchlistData.Tickers {
		parts := strings.SplitN(alert, " ", 2)
		inequality := parts[0]
		price := parts[1]

		row := fmt.Sprintf("('%s', '%s', '%s', '%s', %s)",
			uuid.New().String(),
			watchlistID,
			ticker,
			inequality,
			price,
		)

		rows = append(rows, row)
	}

	alertsQuery += strings.Join(rows, ",\n")
	fmt.Println(alertsQuery)

	_, err = tx.Exec(alertsQuery)
	if err != nil {
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil { log.Fatal(err) }
}