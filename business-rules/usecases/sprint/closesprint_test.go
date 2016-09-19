package sprintuc

import (
	"fmt"
	"testing"

	"github.com/moul/go-clean-architecture/app/repositories/sprint/sprintrepomem"
	"github.com/moul/go-clean-architecture/business-rules/usecases/sprint/dto"
	. "github.com/smartystreets/goconvey/convey"
)

func dummyCloseSprintUC() CloseSprint {
	// prepare sprint repo
	repo := sprintrepomem.New()
	repo.Add(sprintStub1)
	repo.Add(sprintStub2)

	// prepare usecase
	uc := NewCloseSprint()
	uc.SetSprintGateway(repo)
	uc.SetCloseSprintResponseBuilder(sprintucdto.CloseSprintResponseBuilderDTO{})
	return uc
}

func TestCloseSprint(t *testing.T) {
	Convey("Testing CloseSprint", t, FailureContinues, func() {
		// prepare
		uc := dummyCloseSprintUC()
		req := sprintucdto.CloseSprintRequestBuilderDTO{}.Create().WithSprintID(sprintStub2.GetID()).Build()

		// execute usecase
		resp, err := uc.Execute(req)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeNil)
		So(resp.GetClosedIssuesCount(), ShouldEqual, 1)
		So(resp.GetAverageClosedIssues(), ShouldEqual, 1.5)
		So(resp.GetSprintID(), ShouldEqual, sprintStub2.GetID())

		actualSprint, err := uc.gw.Find(sprintStub2.GetID())
		issue := actualSprint.GetIssues()
		So(err, ShouldBeNil)
		So(actualSprint.IsClosed(), ShouldBeTrue)
		So(len(actualSprint.GetIssues()), ShouldEqual, 1)
		So(actualSprint.GetIssues()[0].IsClosed(), ShouldBeTrue)
	})
}
