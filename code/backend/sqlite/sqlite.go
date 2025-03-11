package sqlite

import (
	"database/sql"
	"fmt"
	"net/mail"
	"time"

	"backend/sqlite/errors"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

// Struct to carry the information of a user
type User struct {
	Username    string
	Birthdate   time.Time
	Date_joined time.Time
	Email       string
}

// "Method" to initialise a a new user
func (u *User) InitializeUser(username string, email string, birthdate_string string) (*User, error) {
	birthdate, err := time.Parse("2006-01-02", birthdate_string)
	if err != nil {
		return nil, err
	}

	// Set the date_joined as the current time
	date_joined := time.Now()

	// Check if the email is valid
	_, err = mail.ParseAddress(email)
	if err != nil {
		// Case when the email was not valid
		return nil, err
	}

	// Initialize the user struct with the provided fields
	u.Username = username
	u.Birthdate = birthdate
	u.Date_joined = date_joined
	u.Email = email

	// Return the pointer to the user struct
	return u, nil
}

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

// Function to check if a user exists in the db. Returns a user struct if the user exists, else nil.
func UserExists(db *sql.DB, username string) (*User, error) {
	var user User

	query := "SELECT username, birthdate, date_joined, email  FROM users WHERE username = ?"

	err := db.QueryRow(query, username).Scan(&user.Username, &user.Birthdate, &user.Date_joined, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			// Case when no user is found (not an error)
			return nil, nil
		}
		// Case when an error occurs
		return nil, err
	}

	return &user, nil
}

// Function to add a user to the db. Returns an error if the user already exists.
func AddUser(db *sql.DB, user User, password string) error {

	query := "INSERT INTO users (username, birthdate, date_joined, email, password) VALUES (?, ?, ?, ?, ?)"

	_, err := db.Exec(query, user.Username, user.Birthdate, user.Date_joined, user.Email, password)
	if err != nil {
		// Case when user already exists
		return &errors.UserAlreadyExistsError{Username: user.Username}
	}

	return nil
}

func ClearDb(db *sql.DB, tableName string) error {
	query := fmt.Sprintf("DELETE FROM %s", tableName)
	_, err := db.Exec(query)
	return err
}
