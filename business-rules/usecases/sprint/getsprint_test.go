package sprintuc

import (
	"testing"

	"github.com/moul/go-clean-architecture/app/repositories/sprint/sprintrepomem"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetSprint(t *testing.T) {
	Convey("Testing GetSprint", t, func() {
		uc := NewGetSprint()
		uc.SetSprintGateway(sprintrepomem.SprintRepositoryMemory{})
		//uc.SetGetSprintResponseAssembler(sprintresp.GetSprintResponseAssembler{})
		// reqBuilder := sprintreq.GetSprintRequest{}
	})
}
