package slackbot

import (
	"fmt"
	"regexp"
)

// ParseUserID parses the user ID from an escaped Slack user ID (e.g. `<@UserID>`)
func ParseUserID(escaped string) (string, error) {
	isMatch, err := regexp.MatchString("^\\<\\@.+\\>$", escaped)
	if err != nil {
		return "", err
	}

	if !isMatch {
		return "", fmt.Errorf("'%s' is not in valid escaped Slack user ID format", escaped)
	}

	return escaped[2 : len(escaped)-1], nil
}
