package sprintreq

import "github.com/moul/cleanarch"

type GetSprintRequest interface {
	cleanarch.UseCaseRequest

	GetSprintID() int
}

type GetSprintRequestBuilder interface {
	cleanarch.UseCaseRequestBuilder

	Create() GetSprintRequestBuilder
	WithSprintID(int) GetSprintRequestBuilder
	Build() GetSprintRequest
}

type CloseSprintRequest interface {
	cleanarch.UseCaseRequest

	GetSprintID() int
}
