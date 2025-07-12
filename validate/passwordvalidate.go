package validate

import (
	"errors"
	"regexp"
)

func CheckPassword(password string) error {
	if len(password) < 12 {
		return errors.New("password must be at least 12 characters long")
	}

	if ok, _ := regexp.MatchString(`[a-z]`, password); !ok {
		return errors.New("password must contain at least one lowercase letter")
	}
	if ok, _ := regexp.MatchString(`[A-Z]`, password); !ok {
		return errors.New("password must contain at least one uppercase letter")
	}
	if ok, _ := regexp.MatchString(`[0-9]`, password); !ok {
		return errors.New("password must contain at least one digit")
	}
	if ok, _ := regexp.MatchString(`[!@#$%^&*()_+={}\[\]:;<>,.?/\\|`+"`"+`~\-]`, password); !ok {
		return errors.New("password must contain at least one special character")
	}

	return nil
}
