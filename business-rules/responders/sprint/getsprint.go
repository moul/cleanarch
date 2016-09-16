package sprintresp

import (
	"time"

	"github.com/moul/cleanarch"
	"github.com/moul/go-clean-architecture/business-rules/entities/sprint"
)

type GetSprintResponse interface {
	cleanarch.UseCaseResponse

	GetCreatedAt() time.Time
	GetEffectiveClosedAt() time.Time
	GetExpectedClosedAt() time.Time
	GetID() int
	GetStatus() string
}

type GetSprintResponseAssembler interface {
	Write(*sprint.Sprint) (GetSprintResponse, error)
}
