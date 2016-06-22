package sprintreq

import "github.com/moul/cleanarch"

type GetSprintRequest interface {
	cleanarch.UseCaseRequest

	GetSprintID() int
}
