package sprintuc

import (
	"testing"

	"github.com/moul/go-clean-architecture/app/repositories/sprint/sprintrepomem"
	"github.com/moul/go-clean-architecture/business-rules/usecases/sprint/dto"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetSprint(t *testing.T) {
	Convey("Testing GetSprint", t, func() {
		uc := NewGetSprint()
		uc.SetSprintGateway(sprintrepomem.SprintRepositoryMemory{})
		uc.SetGetSprintResponseAssembler(sprintucdto.GetSprintResponseAssemblerDTO{})
		// reqBuilder := sprintreq.GetSprintRequest{}
	})
}
