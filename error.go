package slackbot

import "fmt"

type UserInputError error

func NewUserInputErrorf(format string, tokens ...interface{}) UserInputError {
	return UserInputError(fmt.Errorf(format, tokens...))
}
