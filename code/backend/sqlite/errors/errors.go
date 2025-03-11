package errors

import "fmt"

// Error if a user is not found
type UserNotFoundError struct {
	Username string
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("user with Name %s not found", e.Username)
}

// Error if the user already exists when trying to implement a new user
type UserAlreadyExistsError struct {
	Username string
}

func (e *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("user with Name %s already exists", e.Username)
}
