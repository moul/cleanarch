package sprintresp

import "github.com/moul/cleanarch"

type CloseSprintResponse interface {
	cleanarch.UseCaseResponse

	GetAverageClosedIssues() int
	GetClosedIssuesCount() int
	GetSprintID() int
}

type CloseSprintResponseBuilder interface {
	Create() CloseSprintResponseBuilder
	WithAverageClosedIssues(int) CloseSprintResponseBuilder
	WithClosedIssuesCount(int) CloseSprintResponseBuilder
	WithSprintID(int) CloseSprintResponseBuilder
	Build() (CloseSprintResponse, error)
}
