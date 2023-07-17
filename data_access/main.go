// https://go.dev/doc/tutorial/database-access
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// Non-`database/sql` compatible custom Postgres-only bindings for
	// Features/Performance.
	// "github.com/jackc/pgx/v5"
	//
	// `database/sql`-compliant library.
	_ "github.com/jackc/pgx/v5/stdlib"
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
	// For the `database/sql` version, which can be swapped to other DB types:
	// https://github.com/jackc/pgx/wiki/Getting-started-with-pgx-through-database-sql
	db, err = sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	pgxQueryExample()
}

// Factored out PGX example from main page to use global `db` pointer.
func pgxQueryExample() {
	var title string
	var artist string
	var price float32
	err := db.QueryRow("select title, artist, price from data_access.album where id=$1", 3).Scan(&title, &artist, &price)
	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	log.Printf("Query output: %v - %v, %v", title, artist, price)
}
