package utils

import (
    "database/sql"
    _"github.com/lib/pq"  // Import the PostgreSQL driver
    "fmt"
    "os"
    _"github.com/joho/godotenv/autoload"
)

func PgConnect() *sql.DB { 
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
	return db
}
