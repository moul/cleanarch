package sprintucdto

import "github.com/moul/go-clean-architecture/business-rules/requestors/sprint"

type GetSprintRequestBuilderDTO struct {
	sprintreq.GetSprintRequestBuilder

	request *GetSprintRequestBuilderDTO
}

//func (b *GetSprintRequestBuilderDTO) Create{
