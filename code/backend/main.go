package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func ConnectDB(path_db string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path_db)
	if err != nil {
		fmt.Println(err)
	}

	// Piung db to see if it returns a value
	err = db.Ping()
	if err != nil {
		// Case when db does not connect
		return nil, err
	}

	return db, err
}

func main() {
	db, err := ConnectDB("../database/test_db.db")
	defer db.Close()

	if err != nil {
		fmt.Println("Error when trying to connect to db")
		return
	}
	fmt.Println("Connected to db")
}

// To run go code: go run main.go
// To build go code: go build -o myapp
