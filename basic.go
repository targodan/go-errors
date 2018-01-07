package errors

import "errors"

// New returns an error that formats as the given text.
func New(text string) error {
	return errors.New(text)
}
