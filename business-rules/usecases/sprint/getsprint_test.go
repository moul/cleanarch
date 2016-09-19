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

func TestGetSprint(t *testing.T) {
	Convey("Testing GetSprint", t, func() {
		// prepare dummy objects
		sprintStub1 := sprint.New()
		sprintStub1.SetID(42)

		// prepare sprint repo
		repo := sprintrepomem.New()
		repo.Add(sprintStub1)

		// prepare usecase
		uc := NewGetSprint()
		uc.SetSprintGateway(repo)
		uc.SetGetSprintResponseAssembler(sprintucdto.GetSprintResponseAssemblerDTO{})

		// prepare request builder
		reqBuilder := sprintucdto.GetSprintRequestBuilderDTO{}
		request := reqBuilder.Create().WithSprintID(42).Build()

		// execute usecase
		resp, err := uc.Execute(request)
		So(err, ShouldBeNil)
		So(resp.GetID(), ShouldEqual, sprintStub1.GetID())
	})
}
