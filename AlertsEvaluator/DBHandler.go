package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Triggered_Alert struct {
	alert_ID   string
	user_email string
}

func DB_connect() *sql.DB {

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
		log.Fatal(pingErr)
	}
	return db
}

func DB_getAlertData(update *PriceUpdate) ([]*Triggered_Alert, error) {

	rows, err := database.Query(`
		SELECT 
			a.id AS alert_id,
			u.email
		FROM alerts a
		JOIN watchlists w ON a.watchlist_id = w.id
		JOIN users u ON w.user_id = u.id
		WHERE a.ticker = $1
		AND a.triggered = false
		AND (
			(a.operator = '>=' AND $2 >= a.target_price) OR
			(a.operator = '<=' AND $2 <= a.target_price) OR
			(a.operator = '='  AND $2 =  a.target_price)
		);`,
		update.Ticker,
		update.Price,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var triggeredAlerts []*Triggered_Alert
	var alertID, userEmail string

	for rows.Next() {
		if err := rows.Scan(&alertID, &userEmail); err != nil {
			return nil, err
		}
		triggeredAlerts = append(triggeredAlerts, &Triggered_Alert{
			alert_ID:   alertID,
			user_email: userEmail,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return triggeredAlerts, nil
}
