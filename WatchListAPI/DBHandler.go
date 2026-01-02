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

// test users:
// eb0dcdff-741d-437c-ad64-35b267a91494
// b573d9e3-5f72-47a4-bc4f-2a882fccb3bb

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

func DB_writeWatchlist(userID string, watchlistData Watchlist) error {
	tx, err := database.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback() // safe rollback if commit never happens

	watchlistID := uuid.New().String()

	_, err = tx.Exec(
		"INSERT INTO watchlists (id, user_id, name) VALUES ($1, $2, $3)",
		watchlistID, userID, watchlistData.Name,
	)
	if err != nil {
		return err
	}

	alertsQuery := "INSERT INTO alerts (id, watchlist_id, ticker, operator, target_price)\nValues"
	rows := []string{}
	for _, alert := range watchlistData.Alerts {

		row := fmt.Sprintf("('%s', '%s', '%s', '%s', %.2f)",
			uuid.New().String(),
			watchlistID,
			alert.Ticker,
			alert.Operator,
			alert.Price,
		)

		rows = append(rows, row)
	}

	alertsQuery += strings.Join(rows, ",\n")

	_, err = tx.Exec(alertsQuery)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func DB_deleteWatchlist(watchlistData DeleteWatchlistsRequest_t) error {
	tx, err := database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM alerts WHERE watchlist_id = $1", watchlistData.ID)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM watchlists WHERE id = $1", watchlistData.ID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func DB_getWatchlists(userData GetWatchlistsRequest_t) (map[string]*Watchlist, error) {
	rows, err := database.Query(
		`SELECT watchlists.name, watchlists.id, alerts.id, alerts.ticker, alerts.operator, alerts.target_price
		FROM watchlists JOIN alerts
		ON alerts.watchlist_id = watchlists.id
		WHERE watchlists.user_id = $1;`,
		userData.ID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	watchlists := make(map[string]*Watchlist)
	var watchlistName, watchlistID, alertID, ticker, operator string
	var targetPrice float32

	for rows.Next() {
		if err := rows.Scan(&watchlistName, &watchlistID, &alertID, &ticker, &operator, &targetPrice); err != nil {
			return nil, err
		}

		// Create watchlist entry if not exists
		WL, exists := watchlists[watchlistID]
		if !exists {
			WL = &Watchlist{
				ID:   watchlistID,
				Name: watchlistName,
			}
			watchlists[watchlistID] = WL
		}

		// Add alert to the watchlist
		WL.Alerts = append(WL.Alerts, &Alert{
			ID:       alertID,
			Ticker:   ticker,
			Operator: operator,
			Price:    targetPrice,
		})
	}

	return watchlists, nil
}
