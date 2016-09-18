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
		// prepare sprint repo
		repo := sprintrepomem.SprintRepositoryMemory{}
		dummySprint := SprintImplTest{
			id: 42,
		}
		//repo.

		// prepare usecase
		uc := NewGetSprint()
		uc.SetSprintGateway(repo)
		uc.SetGetSprintResponseAssembler(sprintucdto.GetSprintResponseAssemblerDTO{})

		// prepare request builder
		reqBuilder := sprintucdto.GetSprintRequestBuilderDTO{}

		request := reqBuilder.Create().WithSprintID(42).Build()
		resp, err := uc.Execute(request)
		So(err, ShouldBeNil)
		So(resp.GetID(), ShouldEqual, dummySprint.GetID())
	})
}
