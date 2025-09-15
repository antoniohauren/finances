package database

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	slog.Info("connecting to database...")

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPass, dbName, dbPort,
	)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("failed to connect to database %v", err)
	}

	slog.Info("connected to database!")

	return db
}
