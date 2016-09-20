package sprintuc

import (
	"github.com/moul/cleanarch"
	"github.com/moul/go-clean-architecture/business-rules/gateways/sprint"
	"github.com/moul/go-clean-architecture/business-rules/requestors/sprint"
	"github.com/moul/go-clean-architecture/business-rules/responders/sprint"
)

type GetSprint struct {
	cleanarch.UseCase

	gw   sprintgw.SprintGateway
	resp sprintresp.GetSprintResponseAssembler
}

func NewGetSprint() GetSprint {
	return GetSprint{}
}

func (uc *GetSprint) SetSprintGateway(gw sprintgw.SprintGateway) {
	uc.gw = gw
}

func (uc *GetSprint) SetGetSprintResponseAssembler(resp sprintresp.GetSprintResponseAssembler) {
	uc.resp = resp
}

func (uc *GetSprint) Execute(req sprintreq.GetSprintRequest) (sprintresp.GetSprintResponse, error) {
	sprint, err := uc.gw.Find(req.GetSprintID())
	if err != nil {
		return nil, err
	}

	return uc.resp.Write(sprint)
}
