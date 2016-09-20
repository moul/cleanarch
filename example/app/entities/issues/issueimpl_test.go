package issueimpl

import (
	"testing"

	"github.com/moul/go-clean-architecture/business-rules/entities/issues"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_IssueImpl(t *testing.T) {
	Convey("Testing IssueImpl", t, func() {
		inst := IssueImpl{}

		err := inst.Close()
		So(err, ShouldBeNil)

		err = inst.Close()
		So(err, ShouldHaveSameTypeAs, &issue.AlreadyClosedError{})

		err = inst.Close()
		So(err, ShouldHaveSameTypeAs, &issue.AlreadyClosedError{})
	})
}
