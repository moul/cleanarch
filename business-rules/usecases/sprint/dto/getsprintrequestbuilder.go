package sprintucdto

import "github.com/moul/go-clean-architecture/business-rules/requestors/sprint"

type GetSprintRequestBuilderDTO struct {
	sprintreq.GetSprintRequestBuilder

	request *GetSprintRequestDTO
}

type GetSprintRequestDTO struct {
	sprintreq.GetSprintRequest

	id int
}

func (r GetSprintRequestDTO) GetSprintID() int {
	return r.id
}

func (b GetSprintRequestBuilderDTO) Create() GetSprintRequestBuilderDTO {
	b.request = &GetSprintRequestDTO{}
	return b
}

func (b GetSprintRequestBuilderDTO) WithSprintID(id int) GetSprintRequestBuilderDTO {
	b.request.id = id
	return b
}

func (b GetSprintRequestBuilderDTO) Build() *GetSprintRequestDTO {
	return b.request
}
