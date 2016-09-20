package addsprint

import (
	"testing"

	"github.com/moul/cleanarch/example/app/repos/sprints/mem"
	"github.com/moul/cleanarch/example/bizrules/usecases"
	"github.com/moul/cleanarch/example/bizrules/usecases/add_sprint/dto"
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
	uc.SetResponseAssembler(addsprintdto.ResponseAssembler{})
	return uc
}

func TestUseCase(t *testing.T) {
	Convey("Testing UseCase", t, FailureContinues, func() {
		// prepare
		uc := dummyUseCase()
		req := addsprintdto.RequestBuilder{}.Create().Build()

		// execute usecase
		resp, err := uc.Execute(req)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeNil)
		So(resp.GetID(), ShouldEqual, 1)

		actualSprint, err := uc.gw.Find(1)
		So(err, ShouldBeNil)
		So(actualSprint.IsClosed(), ShouldBeFalse)
		So(len(actualSprint.GetIssues()), ShouldEqual, 0)
		So(actualSprint.GetCreatedAt(), ShouldResemble, resp.GetCreatedAt())
	})
}
