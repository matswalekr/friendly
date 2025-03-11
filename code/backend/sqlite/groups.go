package sqlite

import (
	"database/sql"
	"sqlite/errors"
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
	Members *Member_list
}

func GroupExists(db *sql.DB, group_name string) (*Group, error) {
	group := &Group{}
	err := db.QueryRow("SELECT * FROM groups WHERE name = ?", group_name).Scan(&group.Name, &group.Owner)
	if err != nil {
		if err == sql.ErrNoRows {
			// Case when no user is found (not an error)
			return nil, nil
		}
		// Case when an error occurs
		return nil, err
	}
	return group, nil
}

func CreateGroup(db *sql.DB, group_name string, owner_name string) error {
	// Make sure that the user exists
	user, err := UserExists(db, owner_name)
	if err != nil {
		return err
	}
	if user == nil {
		return &errors.UserNotFoundError{Username: owner_name}
	}

	// Check if the group already exists
	group, err := GroupExists(db, group_name)
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
	return nil
}
