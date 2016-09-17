package sprintrepomem

import (
	"fmt"

	"github.com/moul/go-clean-architecture/business-rules/entities/sprint"
	"github.com/moul/go-clean-architecture/business-rules/gateways/sprint"
)

type SprintRepositoryMemory struct {
	sprintgw.SprintGateway
}

func (r SprintRepositoryMemory) Find(int) (*sprint.Sprint, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r SprintRepositoryMemory) FindSprintToClose() (*sprint.Sprint, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r SprintRepositoryMemory) FindAverageClosedIssues() int {
	return 0
}

func (r SprintRepositoryMemory) Update(*sprint.Sprint) error {
	return fmt.Errorf("Not implemented")
}
