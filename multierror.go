package errors

// MultiErrorPrefix will be prepended to the Error() output
// of any MultiError, that actually consists of multiple errors.
// You can overwrite this if you want to.
var MultiErrorPrefix = "Multiple errors occured:\n"

// MultiError is an implementation of error that consists
// of arbitrarily many errors on the same logical level.
type MultiError struct {
	Errors []error
}

// NewMultiError returns an error consisting of multiple errors
// on the same logical level.
// If you provide a MultiError, it will append its errors
// correctly.
func NewMultiError(errors ...error) error {
	me := &MultiError{Errors: make([]error, 0)}
	for _, err := range errors {
		if err != nil {
			merr, isMulti := err.(*MultiError)
			if isMulti {
				me.Errors = append(me.Errors, merr.Errors...)
			} else {
				me.Errors = append(me.Errors, err)
			}
		}
	}
	if len(me.Errors) == 0 {
		return nil
	}
	return me
}

func (e *MultiError) Error() string {
	if len(e.Errors) == 1 {
		return e.Errors[0].Error()
	}

	text := MultiErrorPrefix
	for i, err := range e.Errors {
		text += err.Error()
		if i+1 < len(e.Errors) {
			text += "\n"
		}
	}
	return text
}

// IsMultiError returns true if the given error was a
// MultiError.
func IsMultiError(err error) bool {
	_, ok := err.(*MultiError)
	return ok
}
