package sprintucdto

import "github.com/moul/go-clean-architecture/business-rules/responders/sprint"

type CloseSprintResponseBuilderDTO struct {
	sprintresp.CloseSprintResponseBuilder

	response *CloseSprintResponseDTO
}

func (a CloseSprintResponseBuilderDTO) Create() sprintresp.CloseSprintResponseBuilder {
	a.response = &CloseSprintResponseDTO{}
	return a
}

func (a CloseSprintResponseBuilderDTO) WithAverageClosedIssues(val float64) sprintresp.CloseSprintResponseBuilder {
	a.response.averageClosedIssues = val
	return a
}

func (a CloseSprintResponseBuilderDTO) WithClosedIssuesCount(val int) sprintresp.CloseSprintResponseBuilder {
	a.response.closedIssuesCount = val
	return a
}

func (a CloseSprintResponseBuilderDTO) WithSprintID(val int) sprintresp.CloseSprintResponseBuilder {
	a.response.sprintID = val
	return a
}

func (a CloseSprintResponseBuilderDTO) Build() (sprintresp.CloseSprintResponse, error) {
	return a.response, nil
}
