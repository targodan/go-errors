package make

type multiError struct {
	Errors []error
}

// MultiError returns an error consisting of multiple errors.
// If you give it a MultiError, it will append its sub errors
// correctly.
func MultiError(errors ...error) error {
	me := &multiError{Errors: make([]error, 0)}
	for _, err := range errors {
		if err != nil {
			merr, isMulti := err.(*multiError)
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

func (e *multiError) Error() string {
	if len(e.Errors) == 1 {
		return e.Errors[0].Error()
	}

	text := "Multiple errors occured:\n"
	for i, err := range e.Errors {
		text += err.Error()
		if i+1 < len(e.Errors) {
			text += "\n"
		}
	}
	return text
}
