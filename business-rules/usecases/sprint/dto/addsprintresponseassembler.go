package sprintucdto

import (
	"github.com/moul/go-clean-architecture/business-rules/entities/sprint"
	"github.com/moul/go-clean-architecture/business-rules/responders/sprint"
)

type AddSprintResponseAssemblerDTO struct {
	sprintresp.AddSprintResponseAssembler
}

func (a AddSprintResponseAssemblerDTO) Write(sprint *sprint.Sprint) (sprintresp.AddSprintResponse, error) {
	resp := GetSprintResponseDTO{
		createdAt: sprint.GetCreatedAt(),
		id:        sprint.GetID(),
	}
	return resp, nil
}
