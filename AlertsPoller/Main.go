package main

import (
	"database/sql"
	"fmt"
)

const ENV_FILE = "AlertsPoller.env"

var database *sql.DB

func main() {
	database = DB_connect()
	defer database.Close()
	allAlerts, err := DB_getAlerts()

	if err != nil {
		fmt.Println("error getting alerts")
	}

	for _, alert := range allAlerts {
		fmt.Printf("%s\n", alert.ticker)
	}
}
