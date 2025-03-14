package main_test

import (
	"fmt"

	"sqlite"
)

func main() {

	db, err := sqlite.ConnectDB(sqlite.Test_db_path)
	if err != nil {
		fmt.Println("Error when trying to connect to db")
		return
	}
	defer db.Close()

	sqlite.ClearTestDb(db)

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

	user2, err := User_empty.InitializeUser("user2", "walma697@gmail.com", "1997-06-07")
	if err != nil {
		fmt.Println("Error when trying to initialize user2")
		fmt.Println(err)
		return
	}

	err = sqlite.AddUser(db, *user2, "password")
	if err != nil {
		fmt.Println("Error when trying to add user2")
		fmt.Println(err)
		return
	}

	// Try to search for a user
	user, err := sqlite.GetUser(db, "test_user")
	if err != nil {
		fmt.Println("Error when querrying for user")
		fmt.Println(err)
		return
	}
	if user == nil {
		fmt.Println("No user found")
	}

	err = sqlite.AddGroup(db, "test_group", "test_user")
	if err != nil {
		fmt.Println("Error when trying to create group")
		fmt.Println(err)
		return
	}

	err = sqlite.AddMember(db, "test_group", "user2")
	if err != nil {
		fmt.Println("Error when trying to add member to group")
		fmt.Println(err)
		return
	}

	group, err := sqlite.GetGroup(db, "test_group")
	if err != nil {
		fmt.Println("Error when trying to get group")
		fmt.Println(err)
		return
	}
	group.PrintGroupMembers()

	// Clear the test_db
	sqlite.ClearTestDb(db)
}

// To run go code: go run main.go
// To build go code: go build -o myapp
