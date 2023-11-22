package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Connect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect database")
	}

	//defer db.Close()
	return db
}
