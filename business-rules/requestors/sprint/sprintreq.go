package sprintreq

import "github.com/moul/cleanarch"

// GetSprint

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

// AddSprint
type AddSprintRequest interface {
	cleanarch.UseCaseRequest
}
type AddSprintRequestBuilder interface {
	cleanarch.UseCaseRequestBuilder
	Create() AddSprintRequestBuilder
	Build() AddSprintRequest
}

// CloseSprint
type CloseSprintRequest interface {
	cleanarch.UseCaseRequest
	GetSprintID() int
}
type CloseSprintRequestBuilder interface {
	cleanarch.UseCaseRequestBuilder

	Create() GetSprintRequestBuilder
	WithSprintID(int) GetSprintRequestBuilder
	Build() GetSprintRequest
}
