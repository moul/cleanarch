package sprintresp

import (
	"time"

	"github.com/moul/cleanarch"
	"github.com/moul/go-clean-architecture/business-rules/entities/sprint"
)

type SprintResponse interface {
	cleanarch.UseCaseResponse

	GetCreatedAt() time.Time
	GetEffectiveClosedAt() time.Time
	GetExpectedClosedAt() time.Time
	GetID() int
	GetStatus() string
}

type SprintResponseAssembler interface {
	Write(*sprint.Sprint) (SprintResponse, error)
}
