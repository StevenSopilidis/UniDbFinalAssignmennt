package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")

	if username == "" {
		log.Panicln("DB_USERNAME param must be provided")
	}
	if password == "" {
		log.Panicln("DB_PASSWORD param must be provided")
	}
	if databaseName == "" {
		log.Panicln("DB_NAME param must be provided")

	}
	connectionString := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", username, password, databaseName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panicln(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Panicln(err.Error())
	}
	log.Println("Connected to Database succesffully")
	return db
}
