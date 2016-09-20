package sprintucdto

import "github.com/moul/go-clean-architecture/business-rules/responders/sprint"

type CloseSprintResponseDTO struct {
	sprintresp.CloseSprintResponse

	averageClosedIssues float64
	closedIssuesCount   int
	sprintID            int
}

func (r CloseSprintResponseDTO) GetAverageClosedIssues() float64 { return r.averageClosedIssues }
func (r CloseSprintResponseDTO) GetClosedIssuesCount() int       { return r.closedIssuesCount }
func (r CloseSprintResponseDTO) GetSprintID() int                { return r.sprintID }
