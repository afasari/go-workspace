package model

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func Setup() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("could not load env", err)
	}

	db, err = sql.Open("postgres", fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB")))
	if err != nil {
		fmt.Println("could not connect to db", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("could not ping to db", err)
	}
}
