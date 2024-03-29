// https://go.dev/doc/tutorial/database-access
package main

import (
	"database/sql"
	"errors"
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

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

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

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	// Hard-code ID 2 here to test the query.
	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)
}

// Factored out PGX example from main page to use global `db` pointer.
func pgxQueryExample() {
	var title string
	var artist string
	var price float32
	// `QueryRow` returns a single row.
	err := db.QueryRow("select title, artist, price from data_access.album where id=$1", 3).Scan(&title, &artist, &price)
	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	log.Printf("Query output: %v - %v, %v", title, artist, price)
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	// `Query` returns a list of rows to iterate over.
	rows, err := db.Query("SELECT * FROM data_access.album WHERE artist = $1", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// albumByID queries for the album with the specified ID.
func albumByID(id int64) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	row := db.QueryRow("SELECT * FROM data_access.album WHERE id = $1", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO data_access.album (title, artist, price) VALUES ($1, $2, $3)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("Exec - addAlbum: %v", err)
	}
	// `LastInsertId is not supported by this driver`.
	// See: https://github.com/jackc/pgx/issues/1483 -
	// "LastInsertId is not supported by this driver #1483".
	// Suggested workaround is a `db.QueryRow()` to INSERT
	// + `RETURNING id` on the end of the Query.
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		fmt.Println(err.Error())
		fmt.Println(err == errors.New("LastInsertId is not supported by this driver"))
		x := errors.New("LastInsertId is not supported by this driver")
		fmt.Println(err == x)
		fmt.Println(err.Error() == "LastInsertId is not supported by this driver")
		errors.Is(err, errors.New("LastInsertId is not supported by this driver"))
		if err.Error() == x.Error() {
			fmt.Println(result.RowsAffected())
			return 0, fmt.Errorf("addAlbum: Album added, but failed to return ID. %v", err)
		} else {
			return 0, fmt.Errorf("LastInsertId - addAlbum: %v", err)
		}
	}
	return id, nil
}
