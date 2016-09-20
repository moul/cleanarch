package sprintsmem

import (
	"fmt"

	"github.com/moul/cleanarch/example/bizrules/entities"
	"github.com/moul/cleanarch/example/bizrules/gateways"
)

const maxSprintID = 42

type Repo struct {
	gateways.Sprints

	sprints []entities.Sprint
}

func New() *Repo {
	return &Repo{
		sprints: make([]entities.Sprint, 0),
	}
}

func (r *Repo) New() (*entities.Sprint, error) {
	for i := 1; i < maxSprintID; i++ {
		exists := false
		for _, sprint := range r.sprints {
			if sprint.GetID() == i {
				exists = true
				break
			}
		}
		if !exists {
			newSprint := entities.NewSprint()
			newSprint.SetID(i)
			if err := r.Add(newSprint); err != nil {
				return nil, err
			}
			return newSprint, nil
		}
	}
	return nil, fmt.Errorf("too much sprint in the repo")
}

func (r *Repo) Add(sprint *entities.Sprint) error {
	r.sprints = append(r.sprints, *sprint)
	return nil
}

func (r Repo) Find(id int) (*entities.Sprint, error) {
	for _, sprint := range r.sprints {
		if sprint.GetID() == id {
			return &sprint, nil
		}
	}
	return nil, entities.SprintNotFoundError{}
}

func (r Repo) FindSprintToClose() (*entities.Sprint, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r Repo) FindAverageClosedIssues() float64 {
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

func (r *Repo) Update(updated *entities.Sprint) error {
	for idx, sprint := range r.sprints {
		if sprint.GetID() == updated.GetID() {
			r.sprints[idx] = *updated
			return nil
		}
	}
	return entities.SprintNotFoundError{}
}
