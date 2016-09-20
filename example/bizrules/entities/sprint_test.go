package entities

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Sprint(t *testing.T) {
	Convey("Testing Sprint", t, func() {
		sprint := NewSprint()
		sprint.id = 42
		now := time.Now()
		sprint.expectedClosedAt = now
		sprint.effectiveClosedAt = now
		sprint.createdAt = now

		So(sprint.GetID(), ShouldEqual, 42)
		So(sprint.GetCreatedAt(), ShouldResemble, now)
		So(sprint.GetExpectedClosedAt(), ShouldResemble, now)
		So(sprint.GetEffectiveClosedAt(), ShouldResemble, now)
	})
}

func Test_Sprint_AddIssue(t *testing.T) {
	Convey("Testing Sprint.AddIssue()", t, func() {
		sprint := NewSprint()

		So(sprint.GetIssuesCount(), ShouldEqual, 0)
		So(sprint.GetIssuesCount(), ShouldEqual, 0)
		sprint.AddIssue(NewIssue())
		So(sprint.GetIssuesCount(), ShouldEqual, 1)
		So(sprint.GetIssuesCount(), ShouldEqual, 1)
		sprint.AddIssue(NewIssue())
		So(sprint.GetIssuesCount(), ShouldEqual, 2)
		So(sprint.GetIssuesCount(), ShouldEqual, 2)
	})
}

func Test_Sprint_Close(t *testing.T) {
	Convey("Testing Sprint.Close()", t, func() {
		Convey("without issues", func() {
			sprint := NewSprint()

			So(sprint.IsClosed(), ShouldBeFalse)
			So(sprint.IsClosed(), ShouldBeFalse)

			err := sprint.Close()
			So(err, ShouldBeNil)

			So(sprint.IsClosed(), ShouldBeTrue)
			So(sprint.IsClosed(), ShouldBeTrue)

			err = sprint.Close()
			So(err, ShouldHaveSameTypeAs, &SprintAlreadyClosedError{})
			So(err.Error(), ShouldResemble, "Sprint already closed")
		})

		Convey("with issues", func() {
			sprint := NewSprint()

			inst := NewIssue()
			sprint.AddIssue(inst)

			inst = NewIssue()
			err := inst.SetDone()
			So(err, ShouldBeNil)
			sprint.AddIssue(inst)

			So(sprint.IsClosed(), ShouldBeFalse)
			So(sprint.IsClosed(), ShouldBeFalse)

			err = sprint.Close()
			So(err, ShouldBeNil)

			So(sprint.IsClosed(), ShouldBeTrue)
			So(sprint.IsClosed(), ShouldBeTrue)

			err = sprint.Close()
			So(err, ShouldHaveSameTypeAs, &SprintAlreadyClosedError{})
		})
	})
}

func Example_Sprint() {
	sprint := NewSprint()

	sprint.SetID(42)
	sprint.SetExpectedClosedAt(time.Now())
	sprint.SetEffectiveClosedAt(time.Now())

	fmt.Println(sprint.GetID())

	fmt.Println(len(sprint.GetIssues()))
	fmt.Println(sprint.GetIssuesCount())

	fmt.Println(sprint.GetStatus() == SprintOpen)
	fmt.Println(sprint.IsClosed())

	fmt.Println(sprint.Close())

	fmt.Println(sprint.GetStatus())
	fmt.Println(sprint.IsClosed())

	fmt.Println(len(sprint.GetIssues()))
	fmt.Println(sprint.GetIssuesCount())

	// Output:
	// 42
	// 0
	// 0
	// true
	// false
	// <nil>
	// CLOSED
	// true
	// 0
	// 0
}
