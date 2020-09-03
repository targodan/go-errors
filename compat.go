package errors

import (
	"errors"
	"fmt"
)

// Errorf calls fmt.Errorf.
func Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

// Unwrap calls standard API errors.Unwrap.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// As calls standard API errors.As.
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// Is calls standard API errors.Is.
func Is(err, target error) bool {
	return errors.Is(err, target)
}
