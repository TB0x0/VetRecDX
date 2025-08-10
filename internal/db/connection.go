// DB Connections

package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func connectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	openconns := os.Getenv("DB_OPEN_CONNS")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname,
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		fmt.Printf("DB ping failure: %v", err)
	}

	openconns_i, cverr := strconv.Atoi(openconns)

	if cverr != nil {
		fmt.Printf("Environment variable conversion error: %v\n", cverr)
	}

	// After successful connect define DB settings from env
	DB.SetMaxOpenConns(openconns_i)
}
