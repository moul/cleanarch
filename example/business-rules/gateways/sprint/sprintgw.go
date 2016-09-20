package sprintgw

import "github.com/moul/go-clean-architecture/business-rules/entities/sprint"

// SprintGateway is the gateway to the Sprint entity.
type SprintGateway interface {
	Add(*sprint.Sprint) error
	New() (*sprint.Sprint, error)
	Find(int) (*sprint.Sprint, error)
	FindSprintToClose() (*sprint.Sprint, error)
	FindAverageClosedIssues() float64
	Update(*sprint.Sprint) error
}
