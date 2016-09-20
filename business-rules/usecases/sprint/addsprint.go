package sprintuc

import (
	"github.com/moul/cleanarch"
	"github.com/moul/go-clean-architecture/business-rules/gateways/sprint"
	"github.com/moul/go-clean-architecture/business-rules/requestors/sprint"
	"github.com/moul/go-clean-architecture/business-rules/responders/sprint"
)

type AddSprint struct {
	cleanarch.UseCase

	gw   sprintgw.SprintGateway
	resp sprintresp.AddSprintResponseAssembler
}

func NewAddSprint() AddSprint {
	return AddSprint{}
}

func (uc *AddSprint) SetSprintGateway(gw sprintgw.SprintGateway) {
	uc.gw = gw
}

func (uc *AddSprint) SetGetSprintResponseAssembler(resp sprintresp.AddSprintResponseAssembler) {
	uc.resp = resp
}

func (uc *AddSprint) Execute(req sprintreq.AddSprintRequest) (sprintresp.AddSprintResponse, error) {
	newSprint, err := uc.gw.New()
	if err != nil {
		return nil, err
	}

	return uc.resp.Write(newSprint)
}
