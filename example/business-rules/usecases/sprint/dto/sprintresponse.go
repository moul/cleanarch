package sprintucdto

import (
	"time"

	"github.com/moul/go-clean-architecture/business-rules/responders/sprint"
)

type GetSprintResponseDTO struct {
	sprintresp.GetSprintResponse

	createdAt         time.Time
	effectiveClosedAt time.Time
	expectedClosedAt  time.Time
	id                int
	status            string
}

func (r GetSprintResponseDTO) GetCreatedAt() time.Time         { return r.createdAt }
func (r GetSprintResponseDTO) GetEffectiveClosedAt() time.Time { return r.effectiveClosedAt }
func (r GetSprintResponseDTO) GetExpectedClosedAt() time.Time  { return r.expectedClosedAt }
func (r GetSprintResponseDTO) GetID() int                      { return r.id }
func (r GetSprintResponseDTO) GetStatus() string               { return r.status }
