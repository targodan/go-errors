package errors

import (
	"fmt"
	"strings"
)

// HierarchicalErrorLevelSeparator separates the error levels
// in HierarchicalErrors.
var HierarchicalErrorLevelSeparator = "\nPrevious error:\n"

// HierarchicalErrorIndent is prepended to any line in a sub
// of a HierarchicalError.
var HierarchicalErrorIndent = "\t"

// HierarchicalError represents a hierarchical combination
// of errors. This is mostly meant to make errors more user
// friendly while maintaining the original information.
//
//     file, err := os.Open(filename)
//     if err != nil {
//         return errors.Wrapf("could not open file with name \"%s\"", err, filename)
//     }
type HierarchicalError struct {
	TopError error
	SubError error
}

func convert(err error) *HierarchicalError {
	nErr, ok := err.(*HierarchicalError)
	if ok {
		return nErr
	}
	return &HierarchicalError{
		TopError: err,
		SubError: nil,
	}
}

func (e *HierarchicalError) append(err error) {
	if e.SubError == nil {
		e.SubError = err
	} else {
		nSubErr := convert(e.SubError)
		nSubErr.append(err)
		e.SubError = nSubErr
	}
}

func (e *HierarchicalError) Error() string {
	msg := e.TopError.Error()
	if e.SubError != nil {
		msg += HierarchicalErrorLevelSeparator
		lines := strings.Split(e.SubError.Error(), "\n")
		msg += HierarchicalErrorIndent + strings.Join(lines, "\n"+HierarchicalErrorIndent)
	}
	return msg
}

// Wrap creates a new HierarchicalError with the given message and
// a sub error.
func Wrap(newMsg string, subErr error) error {
	return &HierarchicalError{
		TopError: New(newMsg),
		SubError: subErr,
	}
}

// Wrapf creates a new HierarchicalError with the given message and
// a sub error. See fmt.Printf for formatting.
func Wrapf(newMsg string, subErr error, args ...interface{}) error {
	return Wrap(fmt.Sprintf(newMsg, args...), subErr)
}

// WrapErr creates a new HierarchicalError with the same message as the given error and a sub error.
func WrapErr(newErr, subErr error) error {
	nNewErr := convert(newErr)

	nNewErr.append(subErr)

	return nNewErr
}
