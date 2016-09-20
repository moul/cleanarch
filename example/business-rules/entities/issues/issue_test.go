package issue

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Issue(t *testing.T) {
	Convey("Testing Issue", t, func() {
		issue := New()

		now := time.Now()
		issue.id = 42
		issue.title = "Issue 42"
		issue.description = "A dummy issue"
		issue.createdAt = now
		issue.closedAt = now
		issue.doneAt = now

		So(issue.GetID(), ShouldEqual, 42)
		So(issue.GetTitle(), ShouldEqual, "Issue 42")
		So(issue.GetDescription(), ShouldEqual, "A dummy issue")
		So(issue.GetCreatedAt(), ShouldResemble, now)
		So(issue.GetClosedAt(), ShouldResemble, now)
		So(issue.GetDoneAt(), ShouldResemble, now)
	})
}

func Test_IssueSetDone(t *testing.T) {
	Convey("Testing Issue.SetDone()", t, func() {
		issue := New()

		So(issue.GetStatus(), ShouldEqual, "")

		So(issue.SetDone(), ShouldBeNil)
		So(issue.GetStatus(), ShouldEqual, Done)

		err := issue.SetDone()
		So(err, ShouldHaveSameTypeAs, &AlreadyDoneError{})
		So(issue.GetStatus(), ShouldEqual, Done)
		So(err.Error(), ShouldEqual, "Issue already done")

		err = issue.SetDone()
		So(err, ShouldHaveSameTypeAs, &AlreadyDoneError{})
		So(issue.GetStatus(), ShouldEqual, Done)
		So(err.Error(), ShouldEqual, "Issue already done")
	})
}

func Test_IssueClose(t *testing.T) {
	Convey("Testing Issue.Close()", t, func() {
		issue := New()

		So(issue.GetStatus(), ShouldEqual, "")

		So(issue.Close(), ShouldBeNil)
		So(issue.GetStatus(), ShouldEqual, Closed)

		err := issue.Close()
		So(err, ShouldHaveSameTypeAs, &AlreadyClosedError{})
		So(err.Error(), ShouldEqual, "Issue already closed")
		So(issue.GetStatus(), ShouldEqual, Closed)

		err = issue.Close()
		So(err, ShouldHaveSameTypeAs, &AlreadyClosedError{})
		So(err.Error(), ShouldEqual, "Issue already closed")
		So(issue.GetStatus(), ShouldEqual, Closed)
	})
}

func Example_Issue() {
	issue := New()
	issue.SetID(42)
	issue.SetTitle("Issue 42")
	issue.SetStatus(Open)
	issue.SetCreatedAt(time.Now())
	issue.SetDescription("A dummy issue")
	issue.SetClosedAt(time.Now())
	issue.SetDoneAt(time.Now())

	fmt.Println(issue.GetID())
	fmt.Println(issue.GetTitle())
	fmt.Println(issue.GetDescription())

	fmt.Println(issue.GetStatus() == Open)
	fmt.Println(issue.IsDone())
	fmt.Println(issue.IsClosed())

	fmt.Println(issue.Close())

	fmt.Println(issue.GetStatus())
	fmt.Println(issue.IsDone())
	fmt.Println(issue.IsClosed())

	// Output:
	// 42
	// Issue 42
	// A dummy issue
	// true
	// false
	// false
	// <nil>
	// CLOSED
	// false
	// true
}
