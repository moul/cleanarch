package sprintuc

import (
	"github.com/moul/cleanarch"
	"github.com/moul/go-clean-architecture/business-rules/gateways/sprint"
	"github.com/moul/go-clean-architecture/business-rules/requestors/sprint"
	"github.com/moul/go-clean-architecture/business-rules/responders/sprint"
)

type SprintUseCase struct {
	cleanarch.UseCase

	sprintgw   sprintgw.SprintGateway
	sprintresp sprintresp.SprintResponseAssembler
}

func (uc *SprintUseCase) Execute(req sprintreq.GetSprintRequest) (cleanarch.UseCaseResponse, error) {
	sprint, err := uc.sprintgw.Find(req.GetSprintID())
	if err != nil {
		return nil, err
	}

	return uc.sprintresp.Write(sprint)
}
