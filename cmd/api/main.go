package main

import (
	"database/sql"
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/antoniohauren/finances/database"
	"github.com/antoniohauren/finances/internal/handlers"
	"github.com/antoniohauren/finances/internal/repositories"
	"github.com/antoniohauren/finances/internal/services"
	"github.com/joho/godotenv"
)

var db *sql.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("no .env found")
	}

	db = database.ConnectDB()

	database.MigrateAll(db)
}

func main() {
	defer db.Close()

	repos := repositories.New(db)
	services := services.New(repos)
	app := handlers.New(services)

	appPort := os.Getenv("APP_PORT")
	port, err := strconv.Atoi(appPort)

	if err != nil {
		log.Fatal("please check APP_PORT, it should be a number")
	}

	slog.Info("running on", "port", port)

	app.Listen(port)

}
