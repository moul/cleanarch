package sprintreposql

import (
	"fmt"

	"github.com/moul/go-clean-architecture/business-rules/entities/sprint"
	"github.com/moul/go-clean-architecture/business-rules/gateways/sprint"
)

type SprintRepositorySQL struct {
	sprintgw.SprintGateway
}

func (r *SprintRepository) Find(int) (*sprint.Sprint, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *SprintRepository) FindSprintToClose() (*sprint.Sprint, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *SprintRepository) FindAverageClosedIssues() int {
	return 0
}

func (r *SprintRepository) Update(*sprint.Sprint) error {
	return fmt.Errorf("Not implemented")
}
