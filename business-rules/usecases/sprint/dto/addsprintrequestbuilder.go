package sprintucdto

import "github.com/moul/go-clean-architecture/business-rules/requestors/sprint"

type AddSprintRequestBuilderDTO struct {
	sprintreq.AddSprintRequestBuilder

	request *AddSprintRequestDTO
}

type AddSprintRequestDTO struct {
	sprintreq.AddSprintRequest

	id int
}

func (r AddSprintRequestDTO) GetSprintID() int {
	return r.id
}

func (b AddSprintRequestBuilderDTO) Create() sprintreq.AddSprintRequestBuilder {
	b.request = &AddSprintRequestDTO{}
	return b
}

func (b AddSprintRequestBuilderDTO) Build() sprintreq.AddSprintRequest {
	return b.request
}
