package sprintrepomem

import (
	"fmt"

	"github.com/moul/go-clean-architecture/business-rules/entities/sprint"
	"github.com/moul/go-clean-architecture/business-rules/gateways/sprint"
)

const maxSprintID = 42

type SprintRepositoryMemory struct {
	sprintgw.SprintGateway

	sprints []sprint.Sprint
}

func New() *SprintRepositoryMemory {
	return &SprintRepositoryMemory{
		sprints: make([]sprint.Sprint, 0),
	}
}

func (r *SprintRepositoryMemory) New() (*sprint.Sprint, error) {
	for i := 1; i < maxSprintID; i++ {
		exists := false
		for _, sprint := range r.sprints {
			if sprint.GetID() == i {
				exists = true
				break
			}
		}
		if !exists {
			newSprint := sprint.New()
			newSprint.SetID(i)
			if err := r.Add(newSprint); err != nil {
				return nil, err
			}
			return newSprint, nil
		}
	}
	return nil, fmt.Errorf("too much sprint in the repo")
}

func (r *SprintRepositoryMemory) Add(sprint *sprint.Sprint) error {
	r.sprints = append(r.sprints, *sprint)
	return nil
}

func (r SprintRepositoryMemory) Find(id int) (*sprint.Sprint, error) {
	for _, sprint := range r.sprints {
		if sprint.GetID() == id {
			return &sprint, nil
		}
	}
	return nil, sprint.NotFoundError{}
}

func (r SprintRepositoryMemory) FindSprintToClose() (*sprint.Sprint, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r SprintRepositoryMemory) FindAverageClosedIssues() float64 {
	sprintsCount := 0
	issuesCount := 0
	for _, sprint := range r.sprints {
		sprintsCount++
		issuesCount += len(sprint.GetIssues())
	}
	if sprintsCount > 0 {
		return float64(issuesCount) / float64(sprintsCount)
	}
	return float64(0)
}

func (r *SprintRepositoryMemory) Update(updated *sprint.Sprint) error {
	for idx, sprint := range r.sprints {
		if sprint.GetID() == updated.GetID() {
			r.sprints[idx] = *updated
			return nil
		}
	}
	return sprint.NotFoundError{}
}
