package sprintucdto

import "github.com/moul/go-clean-architecture/business-rules/requestors/sprint"

type CloseSprintRequestBuilderDTO struct {
	sprintreq.CloseSprintRequestBuilder

	request *CloseSprintRequestDTO
}

type CloseSprintRequestDTO struct {
	sprintreq.CloseSprintRequest

	id int
}

func (r CloseSprintRequestDTO) GetSprintID() int {
	return r.id
}

func (b CloseSprintRequestBuilderDTO) Create() CloseSprintRequestBuilderDTO {
	b.request = &CloseSprintRequestDTO{}
	return b
}

func (b CloseSprintRequestBuilderDTO) WithSprintID(id int) CloseSprintRequestBuilderDTO {
	b.request.id = id
	return b
}

func (b CloseSprintRequestBuilderDTO) Build() *CloseSprintRequestDTO {
	return b.request
}
