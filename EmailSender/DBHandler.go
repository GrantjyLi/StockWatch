package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

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

func DB_AlertTriggered(alert *Triggered_Alert) (bool, error) {

	_, err := database.Exec(
		`update alerts
		set 
			triggered_at = Now(),
			triggered = true
		where id = $1`,
		alert.Alert_ID,
	)

	if err != nil {
		return false, err
	}
	return true, nil

}
