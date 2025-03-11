package errors

import (
	"fmt"
)

type UserAlreadyExistsError struct {
	Username string
}

// Implement the Error method for the custom error type
func (e *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("User with username '%s' already exists", e.Username)
}

type UserNotFoundError struct {
	Username string
}

// Implement the Error method for the custom error type
func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("User with username '%s' not found", e.Username)
}

// Define the custom error type
type GroupAlreadyExistsError struct {
	GroupName string
}

// Implement the Error method for the custom error type
func (e *GroupAlreadyExistsError) Error() string {
	return fmt.Sprintf("Group with name '%s' already exists", e.GroupName)
}

type GroupDoesNotExist struct {
	GroupName string
}

// Implement the Error method for GroupDoesNotExist
func (e *GroupDoesNotExist) Error() string {
	return fmt.Sprintf("Group with name '%s' does not exist", e.GroupName)
}
