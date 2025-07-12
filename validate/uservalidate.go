package validate

import (
	"errors"
	"regexp"
)

func CheckUsername(username string) error {
	minLen := 3
	maxLen := 32

	if len(username) < minLen || len(username) > maxLen {
		return errors.New("username must be between 3 and 32 characters")
	}

	matched, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, username)
	if !matched {
		return errors.New("username can only contain alphanumeric characters, hyphens, and underscores")
	}

	if matched, _ := regexp.MatchString(`^[-_]`, username); matched {
		return errors.New("username cannot start with a hyphen or underscore")
	}
	if matched, _ := regexp.MatchString(`[-_]$`, username); matched {
		return errors.New("username cannot end with a hyphen or underscore")
	}

	return nil
}
