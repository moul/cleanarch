package closesprint

import (
	"testing"

	"github.com/moul/cleanarch/example/app/repos/sprints/mem"
	"github.com/moul/cleanarch/example/bizrules/usecases"
	"github.com/moul/cleanarch/example/bizrules/usecases/close_sprint/dto"
	. "github.com/smartystreets/goconvey/convey"
)

func dummyUseCase() UseCase {
	// prepare sprint repo
	repo := sprintsmem.New()
	repo.Add(usecases.SprintStub1)
	repo.Add(usecases.SprintStub2)

	// prepare usecase
	uc := New()
	uc.SetSprintsGateway(repo)
	uc.SetResponseBuilder(closesprintdto.ResponseBuilder{})
	return uc
}

func TestUseCase(t *testing.T) {
	Convey("Testing UseCase", t, FailureContinues, func() {
		// prepare
		uc := dummyUseCase()
		req := closesprintdto.RequestBuilder{}.Create().WithSprintID(usecases.SprintStub2.GetID()).Build()

		// execute usecase
		resp, err := uc.Execute(req)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeNil)
		So(resp.GetClosedIssuesCount(), ShouldEqual, 1)
		So(resp.GetAverageClosedIssues(), ShouldEqual, 1.5)
		So(resp.GetSprintID(), ShouldEqual, usecases.SprintStub2.GetID())

		actualSprint, err := uc.gw.Find(usecases.SprintStub2.GetID())
		So(err, ShouldBeNil)
		So(actualSprint.IsClosed(), ShouldBeTrue)
		So(len(actualSprint.GetIssues()), ShouldEqual, 1)
		So(actualSprint.GetIssues()[0].IsClosed(), ShouldBeTrue)
	})
}
