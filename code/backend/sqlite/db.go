package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

// Function to connect to a sqlite3 db. Returns a pointer to the db and an error.
func ConnectDB(path_db string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path_db)
	if err != nil {
		fmt.Println(err)
	}

	// Ping db to see if it returns a value
	err = db.Ping()
	if err != nil {
		// Case when db does not connect
		return nil, err
	}

	return db, err
}

// Function to completely clear a table in a DB
func ClearDb(db *sql.DB, tableName string) error {
	// Construct the query string
	query := fmt.Sprintf("DELETE FROM %s", tableName)

	// Execute the query
	_, err := db.Exec(query)
	return err
}

func ClearTestDb(db *sql.DB) {
	// Clear the entire test_db
	_ = ClearDb(db, "users")
	_ = ClearDb(db, "groups")
	_ = ClearDb(db, "group_members")
}
