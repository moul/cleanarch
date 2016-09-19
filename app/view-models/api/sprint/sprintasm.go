package app_vm_api_sprint

import (
	"encoding/json"
	"time"

	"github.com/moul/go-clean-architecture/business-rules/responders/sprint"
)

type GetAssembler struct{}

func NewGetAssembler() *GetAssembler { return &GetAssembler{} }

func (a *GetAssembler) WriteGet(resp sprintresp.GetSprintResponse) Get {
	get := Get{}
	get.CreatedAt = resp.GetCreatedAt()
	get.EffectiveClosedAt = resp.GetEffectiveClosedAt()
	get.ExpectedClosedAt = resp.GetExpectedClosedAt()
	get.Status = resp.GetStatus()
	return get
}

type Get struct {
	CreatedAt         time.Time `json:"created-at"`
	EffectiveClosedAt time.Time `json:"effective-closed-at"`
	ExpectedClosedAt  time.Time `json:"expected-closed-at"`
	Status            string    `json:"status"`
}

func (g *Get) Serialize() []byte {
	out, _ := json.Marshal(*g)
	return out
}
