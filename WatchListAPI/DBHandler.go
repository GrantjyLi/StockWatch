package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

const ENV_FILE = "WatchListAPI.env"

// TODO: change this later
var userID = "eb0dcdff-741d-437c-ad64-35b267a91494"

//var userID = "b573d9e3-5f72-47a4-bc4f-2a882fccb3bb"

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

func DB_writeWatchlist(watchlistData Watchlist) {
	tx, err := database.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback() // safe rollback if commit never happens

	watchlistID := uuid.New().String()

	watchlistQuery := fmt.Sprintf("INSERT INTO watchlists (id, user_id, name)\nVALUES ('%s', '%s', '%s');",
		watchlistID,
		userID,
		watchlistData.Name,
	)

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

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func DB_getWatchlists(userData GetWatchlistsRequest) (map[string]*Watchlist, error) {
	rows, err := database.Query(
		`SELECT watchlists.name, watchlists.id, alerts.ticker, alerts.operator, alerts.target_price
		FROM watchlists JOIN alerts
		ON alerts.watchlist_id = watchlists.id
		WHERE watchlists.user_id = $1;`,
		userData.ID,
	)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	watchlists := make(map[string]*Watchlist)
	var watchlistName, watchlistID, ticker, operator string
	var targetPrice float64

	for rows.Next() {
		if err := rows.Scan(&watchlistName, &watchlistID, &ticker, &operator, &targetPrice); err != nil {
			log.Fatal(err)
		}

		// Create watchlist entry if not exists
		WL, exists := watchlists[watchlistID]
		if !exists {
			WL = &Watchlist{
				Name:    watchlistName,
				Tickers: make(map[string]string),
			}
			watchlists[watchlistID] = WL
		}

		// Add alert to the watchlist
		WL.Tickers[ticker] = fmt.Sprintf("%s %.2f", operator, targetPrice)
	}

	// for id, watchlist := range watchlists {
	// 	fmt.Printf("%s:\n", id)
	// 	fmt.Printf("    %s\n", watchlist.Name)
	// 	for ticker, condition := range watchlist.Tickers {
	// 		fmt.Printf("    %s:%s\n", ticker, condition)
	// 	}
	// }

	return watchlists, nil
}
