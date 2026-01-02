package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type Alert struct {
	ticker string
}

var conditions = []string{
	fmt.Sprintf("alerts.triggered = %s", "FALSE"),
}

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

func DB_getAlerts() ([]*Alert, error) {
	alertsQuery := "SELECT DISTINCT alerts.ticker FROM alerts"

	if len(conditions) > 0 {
		alertsQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := database.Query(alertsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alerts []*Alert
	var alertTicker string

	for rows.Next() {
		if err := rows.Scan(&alertTicker); err != nil {
			return nil, err
		}
		alerts = append(alerts, &Alert{
			ticker: alertTicker,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return alerts, nil
}
