package sprintuc

import (
	"github.com/moul/cleanarch"
	"github.com/moul/go-clean-architecture/business-rules/gateways/sprint"
	"github.com/moul/go-clean-architecture/business-rules/requestors/sprint"
	"github.com/moul/go-clean-architecture/business-rules/responders/sprint"
)

type CloseSprint struct {
	cleanarch.UseCase

	gw   sprintgw.SprintGateway
	resp sprintresp.CloseSprintResponseBuilder
}

func (uc *CloseSprint) Execute(req sprintreq.CloseSprintRequest) (cleanarch.UseCaseResponse, error) {
	sprint, err := uc.gw.Find(req.GetSprintID())
	if err != nil {
		return nil, err
	}

	if err := sprint.Close(); err != nil {
		return nil, err
	}

	if err := uc.gw.Update(sprint); err != nil {
		return nil, err
	}

	return uc.resp.
		Create().
		WithAverageClosedIssues(uc.gw.FindAverageClosedIssues()).
		WithClosedIssuesCount(sprint.GetIssuesCount()).
		WithSprintID(sprint.GetID()).
		Build()
}
