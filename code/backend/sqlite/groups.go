package sqlite

import (
	"database/sql"
	"fmt"
	"sqlite/errors"
	"time"
)

type Node_member struct {
	Next     *Node_member
	Username string
}

type Member_list struct {
	head *Node_member
}

func (list *Member_list) Append(username string) {
	newNode := &Node_member{Username: username}
	if list.head == nil {
		list.head = newNode
		return
	} else {
		current := list.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
		return
	}
}

func (list *Member_list) Remove(username string) {
	if list.head == nil {
		return
	}
	if list.head.Username == username {
		list.head = list.head.Next
		return
	}
	current := list.head
	for current.Next != nil {
		if current.Next.Username == username {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

type Group struct {
	Name    string
	Owner   string
	Id      int
	Members Member_list
}

func (g *Group) PrintGroupMembers() {
	fmt.Println("Group members:")
	if g.Members.head == nil {
		return
	}
	current := g.Members.head
	for current != nil {
		println(current.Username)
		current = current.Next
	}
}

func GetGroup(db *sql.DB, group_name string) (*Group, error) {
	var err error
	group := &Group{}

	err = db.QueryRow("SELECT name, group_id FROM groups WHERE name = ?", group_name).Scan(&group.Name, &group.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			// Case when no group is found (not an error)
			return nil, nil
		}
		// Case when an error occurs
		return nil, err
	}

	// Get the Owner
	err = db.QueryRow("SELECT username FROM users WHERE user_id = (SELECT owner_id FROM groups WHERE name = ?)", group_name).Scan(&group.Owner)
	if err != nil {
		return nil, err
	}

	// Get the members of the group
	rows, err := db.Query("SELECT username FROM users WHERE user_id = (SELECT user_id FROM group_members WHERE group_id = ?)", group.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	group.Members = Member_list{}

	var username string
	for rows.Next() {
		if err := rows.Scan(&username); err != nil {
			return nil, err
		}
		// Add the user to the list of users
		group.Members.Append(username)
	}

	// Check if there were any errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Return the group
	return group, nil
}

func AddGroup(db *sql.DB, group_name string, owner_name string) error {
	// Make sure that the user exists
	user, err := GetUser(db, owner_name)
	if err != nil {
		return err
	}
	if user == nil {
		return &errors.UserNotFoundError{Username: owner_name}
	}

	// Check if the group already exists
	group, err := GetGroup(db, group_name)
	if err != nil {
		return err
	}
	if group != nil {
		return &errors.GroupAlreadyExistsError{GroupName: group_name}
	}

	// If user and group are valid, create the group
	_, err = db.Exec("INSERT INTO groups (name, owner_id) VALUES (?, (SELECT user_id FROM users WHERE username = ?))", group_name, owner_name)
	if err != nil {
		return err
	}

	// Add the owner to the group_members
	err = AddMember(db, group_name, owner_name)
	if err != nil {
		// If this did not work, delete the group
		db.Exec("DELETE FROM groups WHERE name = ?", group_name)
		return err
	}

	return nil
}

func AddMember(db *sql.DB, group_name string, member_name string) error {
	// Make sure that the user exists
	user, err := GetUser(db, member_name)
	if err != nil {
		return err
	}
	if user == nil {
		return &errors.UserNotFoundError{Username: member_name}
	}

	// Check if the group already exists
	group, err := GetGroup(db, group_name)
	if err != nil {
		return err
	}
	if group == nil {
		return &errors.GroupDoesNotExist{GroupName: group_name}
	}

	// Get todays date
	today := time.Now().Format("2006-01-02")

	_, err = db.Exec("INSERT INTO group_members (group_id, member, date_joined) VALUES ((SELECT group_id FROM groups WHERE name = ?), (SELECT user_id FROM users WHERE username = ?), ?)", group_name, member_name, today)
	if err != nil {
		return err
	}

	return nil
}
