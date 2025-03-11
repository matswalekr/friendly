package main

import (
	"fmt"

	"backend/sqlite"
)

// Globally scope constants
const test_db_path string = "../database/test_db.db"
const main_db_path string = "../database/database.db"

func main() {
	db, err := sqlite.ConnectDB(test_db_path)

	if err != nil {
		fmt.Println("Error when trying to connect to db")
		return
	}
	defer db.Close()

	fmt.Println("Connected to db")

	var User_empty sqlite.User

	new_user, err := User_empty.InitializeUser("test_user", "walma697@gmail.com", "1997-06-07")
	if err != nil {
		fmt.Println("Error when trying to initialize user")
		fmt.Println(err)
		return
	}

	err = sqlite.AddUser(db, *new_user, "password")
	if err != nil {
		fmt.Println("Error when trying to add user")
		fmt.Println(err)
		return
	}

	// Try to search for a user
	user, err := sqlite.UserExists(db, "test_user")
	if err != nil {
		fmt.Println("Error when querrying for user")
		fmt.Println(err)
		return
	}

	if user == nil {
		fmt.Println("No user found")
	} else {
		fmt.Printf("User %s found\n", user.Username)
	}

	// Clear the test_db
	err = sqlite.ClearDb(db, "users")
	if err == nil {
		fmt.Printf("Db cleared successfully\n")
	}
}

// To run go code: go run main.go
// To build go code: go build -o myapp
