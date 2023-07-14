package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	// Read `.env` file to populate ENV.
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	pingErr := db.Ping(context.Background())
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	var title string
	var artist string
	var price float32
	err = db.QueryRow(context.Background(), "select title, artist, price from data_access.album where id=$1", 3).Scan(&title, &artist, &price)
	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	log.Printf("Query output: %v - %v, %v", title, artist, price)
}
