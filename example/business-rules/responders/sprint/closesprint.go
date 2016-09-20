package sprintresp

import "github.com/moul/cleanarch"

type CloseSprintResponse interface {
	cleanarch.UseCaseResponse

	GetAverageClosedIssues() float64
	GetClosedIssuesCount() int
	GetSprintID() int
}

type CloseSprintResponseBuilder interface {
	Create() CloseSprintResponseBuilder
	WithAverageClosedIssues(float64) CloseSprintResponseBuilder
	WithClosedIssuesCount(int) CloseSprintResponseBuilder
	WithSprintID(int) CloseSprintResponseBuilder
	Build() (CloseSprintResponse, error)
}
