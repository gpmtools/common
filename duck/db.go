package duck

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/marcboeker/go-duckdb"
)

func NewDB() *sql.DB {
	db, err := sql.Open("duckdb", "")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE people (id INTEGER, name VARCHAR)`)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	_, err = db.Exec(`INSERT INTO people VALUES (42, 'John')`)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var (
		id   int
		name string
	)
	row := db.QueryRow(`SELECT id, name FROM people`)
	err = row.Scan(&id, &name)
	if errors.Is(err, sql.ErrNoRows) {
		log.Println("no rows")
	} else if err != nil {
		log.Fatal(err)
		return nil
	}

	fmt.Printf("id: %d, name: %s\n", id, name)
	return db
}
