package gateways

import "github.com/moul/cleanarch/example/bizrules/entities"

// Sprints is the gateway to the Sprint entity.
type Sprints interface {
	Add(*entities.Sprint) error
	New() (*entities.Sprint, error)
	Find(int) (*entities.Sprint, error)
	FindSprintToClose() (*entities.Sprint, error)
	FindAverageClosedIssues() float64
	Update(*entities.Sprint) error
}
