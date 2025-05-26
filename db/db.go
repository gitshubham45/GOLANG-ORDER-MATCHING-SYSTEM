package db

import (
	"os"
	"log"
	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ")/" + dbName

	var errOpen error
	DB , errOpen = sql.Open("mysql" , dsn)
	if errOpen != nil {
		log.Fatal(errOpen)
	}

	if err := DB.Ping() ; err != nil {
		log.Fatal(err)
	}
}