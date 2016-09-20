package sprintuc

import (
	"testing"

	"github.com/moul/go-clean-architecture/app/repositories/sprint/sprintrepomem"
	"github.com/moul/go-clean-architecture/business-rules/usecases/sprint/dto"
	. "github.com/smartystreets/goconvey/convey"
)

func dummyAddSprintUC() AddSprint {
	// prepare sprint repo
	repo := sprintrepomem.New()
	repo.Add(sprintStub1)
	repo.Add(sprintStub2)

	// prepare usecase
	uc := NewAddSprint()
	uc.SetSprintGateway(repo)
	uc.SetAddSprintResponseAssembler(sprintucdto.AddSprintResponseAssemblerDTO{})
	return uc
}

func TestAddSprint(t *testing.T) {
	Convey("Testing AddSprint", t, FailureContinues, func() {
		// prepare
		uc := dummyAddSprintUC()
		req := sprintucdto.AddSprintRequestBuilderDTO{}.Create().Build()

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
