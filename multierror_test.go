package errors

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewMultiError(t *testing.T) {
	Convey("basic error creation should work", t, func() {
		errs := []error{errors.New("err1"), errors.New("err2"), errors.New("err3")}
		me := NewMultiError(errs...)

		So(me.Errors, ShouldResemble, errs)
	})
	Convey("nil values should be filtered", t, func() {
		errs := []error{nil, errors.New("err1"), nil, errors.New("err3"), nil}
		expErrs := []error{errors.New("err1"), errors.New("err3")}
		me := NewMultiError(errs...)

		So(me.Errors, ShouldResemble, expErrs)
	})
	Convey("the return value of MultiError should be nil", t, func() {
		Convey("if it is presented with only one nil error", func() {
			err := NewMultiError(nil)
			So(err, ShouldBeNil)
		})
		Convey("if it is presented with multiple nil errors", func() {
			err := NewMultiError(nil, nil, nil, nil, nil)
			So(err, ShouldBeNil)
		})
	})
	Convey("other multiErrors should be correctly appended", t, func() {
		innerErrs := []error{errors.New("errA"), errors.New("errB")}
		outerErrs := []error{errors.New("err1"), errors.New("err2"), NewMultiError(innerErrs...), errors.New("err3")}
		expected := []error{errors.New("err1"), errors.New("err2"), errors.New("errA"), errors.New("errB"), errors.New("err3")}

		me := NewMultiError(outerErrs...)

		So(me.Errors, ShouldResemble, expected)
	})
}

func TestMultiErrorString(t *testing.T) {
	Convey("simple errors should stay simple", t, func() {
		simple := errors.New("just a simple error")
		err := NewMultiError(simple)

		So(err.Error(), ShouldEqual, simple.Error())
	})
	Convey("on multiple errors", t, func() {
		errA := errors.New("an error")
		errB := errors.New("another error")
		err := NewMultiError(errA, errB)

		Convey("all errors should be separated by a newline", func() {
			So(err.Error(), ShouldContainSubstring, errA.Error()+"\n"+errB.Error())
		})
		Convey("there should be a prefix", func() {
			So(err.Error(), ShouldStartWith, MultiErrorPrefix)
		})
		Convey("the message should not end on a newline", func() {
			text := err.Error()
			So(text[len(text)-1], ShouldNotEqual, '\n')
		})
	})
}

func TestIsMultiError(t *testing.T) {
	Convey("should be true on MultiError", t, func() {
		Convey("if created manually", func() {
			err := &MultiError{}
			So(IsMultiError(err), ShouldBeTrue)
		})
		Convey("if created via New", func() {
			err := NewMultiError(errors.New("some error"))
			So(IsMultiError(err), ShouldBeTrue)
		})
	})
}
