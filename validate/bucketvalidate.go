// validate/bucket.go
package validate

import (
	"errors"
	"fmt"
	"strings"
)

func BucketValidate(name string, existing []string) error {
	name = strings.ToLower(name)

	if strings.Contains(name, " ") {
		return errors.New("bucket name must not contain spaces")
	}

	for _, b := range existing {
		if strings.ToLower(b) == name {
			return errors.New("bucket already exists")
		}
	}
	return nil
}
func BucketDeleteValidate(name string, existing []string) error {
	name = strings.ToLower(strings.TrimSpace(name))
	for _, b := range existing {
		if strings.ToLower(b) == name {
			return nil // ✅ found, okay to delete
		}
	}
	return fmt.Errorf("bucket '%s' does not exist", name)
}
