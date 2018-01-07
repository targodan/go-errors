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
