package sprintuc

import (
	"testing"

	"github.com/moul/go-clean-architecture/app/repositories/sprint/sprintrepomem"
	"github.com/moul/go-clean-architecture/business-rules/entities/sprint"
	"github.com/moul/go-clean-architecture/business-rules/usecases/sprint/dto"
	. "github.com/smartystreets/goconvey/convey"
)

type SprintImplTest struct {
	sprint.Sprint

	id int
}

var sprintStub1 = sprint.New()

func init() {
	sprintStub1.SetID(42)
}

func getTestingUseCase() GetSprint {
	// prepare sprint repo
	repo := sprintrepomem.New()
	repo.Add(sprintStub1)

	// prepare usecase
	uc := NewGetSprint()
	uc.SetSprintGateway(repo)
	uc.SetGetSprintResponseAssembler(sprintucdto.GetSprintResponseAssemblerDTO{})
	return uc
}

func TestGetSprint(t *testing.T) {
	Convey("Testing GetSprint", t, func() {
		// prepare
		uc := getTestingUseCase()
		req := sprintucdto.GetSprintRequestBuilderDTO{}.Create().WithSprintID(42).Build()

		// execute usecase
		resp, err := uc.Execute(req)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeNil)
		So(resp.GetID(), ShouldEqual, sprintStub1.GetID())
		So(resp.GetStatus(), ShouldEqual, sprintStub1.GetStatus())
		So(resp.GetCreatedAt(), ShouldNotEqual, sprintStub1.GetCreatedAt())
		So(resp.GetCreatedAt(), ShouldResemble, sprintStub1.GetCreatedAt())
		So(resp.GetEffectiveClosedAt(), ShouldNotEqual, sprintStub1.GetEffectiveClosedAt())
		So(resp.GetEffectiveClosedAt(), ShouldResemble, sprintStub1.GetEffectiveClosedAt())
		So(resp.GetExpectedClosedAt(), ShouldNotEqual, sprintStub1.GetExpectedClosedAt())
		So(resp.GetExpectedClosedAt(), ShouldResemble, sprintStub1.GetExpectedClosedAt())
	})

	Convey("Testing NotFound", t, func() {
		// prepare
		uc := getTestingUseCase()
		req := sprintucdto.GetSprintRequestBuilderDTO{}.Create().WithSprintID(43).Build()

		// execute usecase
		resp, err := uc.Execute(req)
		So(err, ShouldNotBeNil)
		So(resp, ShouldBeNil)
		So(err, ShouldResemble, sprint.NotFoundError{})
	})
}
