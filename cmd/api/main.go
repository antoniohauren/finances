package main

import (
	"database/sql"
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/antoniohauren/finances/internal/controllers"
	"github.com/antoniohauren/finances/internal/database"
	"github.com/joho/godotenv"
)

var db *sql.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("no .env found")
	}

	db = database.ConnectDB()
}

func main() {
	defer db.Close()

	app := controllers.New()

	appPort := os.Getenv("APP_PORT")
	port, err := strconv.Atoi(appPort)

	if err != nil {
		log.Fatal("please check APP_PORT, it should be a number")
	}

	slog.Info("running on", "port", port)

	app.Listen(port)

}
