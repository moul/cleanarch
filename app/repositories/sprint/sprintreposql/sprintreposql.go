package sprintreposql

import (
	"fmt"

	"github.com/moul/go-clean-architecture/business-rules/entities/sprint"
	"github.com/moul/go-clean-architecture/business-rules/gateways/sprint"
)

type SprintRepositorySQL struct {
	sprintgw.SprintGateway
}

func (r SprintRepositorySQL) Add(sprint *sprint.Sprint) error {
}

func (r SprintRepositorySQL) Find(int) (*sprint.Sprint, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r SprintRepositorySQL) FindSprintToClose() (*sprint.Sprint, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r SprintRepositorySQL) FindAverageClosedIssues() int {
	return 0
}

func (r SprintRepositorySQL) Update(*sprint.Sprint) error {
	return fmt.Errorf("Not implemented")
}
