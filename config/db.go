package config

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnDB() (*sql.DB, error) {
	godotenv.Load()
	dbStr := os.Getenv("DB_STRING")

	db, err := sql.Open("mysql", dbStr)
	return db, err
}
