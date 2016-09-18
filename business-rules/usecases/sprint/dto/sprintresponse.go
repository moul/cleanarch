package sprintucdto

import (
	"time"

	"github.com/moul/go-clean-architecture/business-rules/responders/sprint"
)

type SprintResponseDTO struct {
	sprintresp.GetSprintResponse

	createdAt         time.Time
	effectiveClosedAt time.Time
	expectedClosedAt  time.Time
	id                int
	status            string
}

func (r *SprintResponseDTO) GetCreatedAt() time.Time         { return r.createdAt }
func (r *SprintResponseDTO) GetEffectiveClosedAt() time.Time { return r.effectiveClosedAt }
func (r *SprintResponseDTO) GetExpectecClosedAt() time.Time  { return r.expectedClosedAt }
func (r *SprintResponseDTO) GetID() int                      { return r.id }
func (r *SprintResponseDTO) GetStatis() string               { return r.status }
