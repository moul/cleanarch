package sprintgw

import "github.com/moul/go-clean-architecture/business-rules/entities/sprint"

// SprintGateway is the gateway to the Sprint entity.
type SprintGateway interface {
	Find(int) (*sprint.Sprint, error)
	FindSprintToClose() (*sprint.Sprint, error)
	FindAverageClosedIssues() int
	Update(*sprint.Sprint) error
}
