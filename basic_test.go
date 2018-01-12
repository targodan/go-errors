package errors

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("New should simply create a new error", t, func() {
		err := New("some err")
		So(err.Error(), ShouldEqual, "some err")
	})
}

func TestNewf(t *testing.T) {
	Convey("Newf without formatting should behave the same as New", t, func() {
		err := Newf("some text")
		So(err.Error(), ShouldEqual, "some text")
	})
	Convey("Newf should format text", t, func() {
		err := Newf("some text with a number %04d", 93)
		So(err.Error(), ShouldEqual, "some text with a number 0093")
	})
}
