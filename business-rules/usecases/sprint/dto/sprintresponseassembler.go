package sprintucdto

import (
	"github.com/moul/go-clean-architecture/business-rules/entities/sprint"
	"github.com/moul/go-clean-architecture/business-rules/responders/sprint"
)

type GetSprintResponseAssemblerDTO struct {
	sprintresp.GetSprintResponseAssembler
}

func (a GetSprintResponseAssemblerDTO) Write(sprint *sprint.Sprint) (sprintresp.GetSprintResponse, error) {
	resp := GetSprintResponseDTO{
		createdAt:         sprint.GetCreatedAt(),
		effectiveClosedAt: sprint.GetEffectiveClosedAt(),
		expectedClosedAt:  sprint.GetExpectedClosedAt(),
		id:                sprint.GetID(),
		status:            sprint.GetStatus(),
	}
	return resp, nil
}
