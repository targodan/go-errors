package errors

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConvert(t *testing.T) {
	Convey("Hierarchical errors should not be converted.", t, func() {
		err := &HierarchicalError{
			TopError: New("top err"),
			SubError: New("sub err"),
		}
		conv := convert(err)

		So(conv, ShouldResemble, err)
		So(conv, ShouldEqual, err)
	})
	Convey("Non-hierarchical errors should be turned into one.", t, func() {
		err := New("basic error")
		conv := convert(err)

		So(conv.TopError, ShouldResemble, err)
		So(conv.TopError, ShouldEqual, err)
		So(conv.SubError, ShouldBeNil)
	})
}

func TestAppend(t *testing.T) {
	Convey("Append should just set the sub error if previously nil", t, func() {
		topErr := New("basic")
		err := &HierarchicalError{
			TopError: topErr,
			SubError: nil,
		}
		subErr := New("appended")
		err.append(subErr)
		exp := &HierarchicalError{
			TopError: topErr,
			SubError: subErr,
		}
		So(err, ShouldResemble, exp)
	})
	Convey("Append should create a new HierarchicalError if sub error is not nil.", t, func() {
		topErr := New("basic")
		subErr := New("basic2")
		err := &HierarchicalError{
			TopError: topErr,
			SubError: subErr,
		}
		nSubErr := New("appended")
		err.append(nSubErr)
		exp := &HierarchicalError{
			TopError: topErr,
			SubError: &HierarchicalError{
				TopError: subErr,
				SubError: nSubErr,
			},
		}
		So(err, ShouldResemble, exp)
	})
}

func TestWrap(t *testing.T) {
	Convey("Wrap should wrap an error in a HierarchicalError.", t, func() {
		text := "some outer text"
		origErr := New("some inner err")
		err := Wrap(text, origErr)
		exp := &HierarchicalError{
			TopError: New(text),
			SubError: origErr,
		}

		So(err, ShouldResemble, exp)
	})
}

func TestWrapErr(t *testing.T) {
	Convey("WrapErr should wrap an error in a HierarchicalError.", t, func() {
		outer := New("outer")
		inner := New("inner")
		err := WrapErr(outer, inner)
		exp := &HierarchicalError{
			TopError: outer,
			SubError: inner,
		}

		So(err, ShouldResemble, exp)
	})
}

func TestWrapf(t *testing.T) {
	Convey("Wrapf should wrap an error in a HierarchicalError.", t, func() {
		text := "some outer text with a number %04d"
		number := 98
		textf := "some outer text with a number 0098"
		origErr := New("some inner err")
		err := Wrapf(text, origErr, number)
		exp := &HierarchicalError{
			TopError: New(textf),
			SubError: origErr,
		}

		So(err, ShouldResemble, exp)
	})
}

func TestError(t *testing.T) {
	Convey("Without a subErr only the basic message should be displayed", t, func() {
		text := "basic error text"
		err := &HierarchicalError{
			TopError: New(text),
			SubError: nil,
		}

		So(err.Error(), ShouldEqual, text)
	})
	Convey("Top and SubErr should be separated properly.", t, func() {
		topText := "top err"
		subText := "sub err"
		err := &HierarchicalError{
			TopError: New(topText),
			SubError: New(subText),
		}

		So(err.Error(), ShouldEqual, topText+HierarchicalErrorLevelSeparator+HierarchicalErrorIndent+subText)
	})
	Convey("All line of the sub error should be indented.", t, func() {
		topText := "top err"
		subText := "sub err\nwith\nmultiple\nlines"
		err := &HierarchicalError{
			TopError: New(topText),
			SubError: New(subText),
		}
		indentedSubText := HierarchicalErrorIndent + strings.Join(strings.Split(subText, "\n"), "\n"+HierarchicalErrorIndent)

		So(err.Error(), ShouldEqual, topText+HierarchicalErrorLevelSeparator+indentedSubText)
	})
}
