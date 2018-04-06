package slackbot

import "fmt"

// UserInputError is a custom error type used when a user sends bad input to a command
type UserInputError struct {
	message string
}

// NewUserInputError creates a new UserInputError with the specified message
func NewUserInputError(message string) *UserInputError {
	return &UserInputError{message}
}

// NewUserInputError creates a new UserInputError with the specified formatted message
func NewUserInputErrorf(format string, tokens ...interface{}) *UserInputError {
	return NewUserInputError(fmt.Sprintf(format, tokens...))
}

// Error is used to satisfy the error interface
func (e *UserInputError) Error() string {
	return e.message
}
