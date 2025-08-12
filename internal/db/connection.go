// DB Connections

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	openconns := os.Getenv("DB_OPEN_CONNS")
	maxidle := os.Getenv("DB_MAX_IDLE")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname,
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("DB open failure: %v", err)
	}

	openconns_i, err := strconv.Atoi(openconns)

	if err != nil {
		log.Panicf("Environment variable conversion error (DB_OPEN_CONNS): %v\n", err)
	}

	maxidle_i, err := strconv.Atoi(maxidle)

	if err != nil {
		log.Panicf("Environment variable conversion error (DB_MAX_IDLE): %v\n", err)
	}

	// After successful connect define DB settings from env
	DB.SetMaxOpenConns(openconns_i)
	DB.SetMaxIdleConns(maxidle_i)
	DB.SetConnMaxLifetime(time.Hour)

	if err := DB.Ping(); err != nil {
		log.Fatalf("DB ping failure: %v \n", err)
	}

	fmt.Println("DB Connected Successfully")
}
