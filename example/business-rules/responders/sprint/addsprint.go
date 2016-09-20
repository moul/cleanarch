package sprintresp

import (
	"time"

	"github.com/moul/cleanarch"
	"github.com/moul/go-clean-architecture/business-rules/entities/sprint"
)

type AddSprintResponse interface {
	cleanarch.UseCaseResponse

	GetCreatedAt() time.Time
	GetID() int
}

type AddSprintResponseAssembler interface {
	Write(*sprint.Sprint) (AddSprintResponse, error)
}
