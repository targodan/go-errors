package errors

import (
	"errors"
	"fmt"
)

// New returns an error that formats as the given text.
func New(text string) error {
	return errors.New(text)
}

// Newf creates a new error with the given message.
// See fmt.Printf and Errorf for formatting.
// Calls standard errors.Errorf, so "%w" is supported.
func Newf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}
