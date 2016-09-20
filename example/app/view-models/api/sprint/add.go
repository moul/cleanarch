package app_vm_api_sprint

import (
	"encoding/json"
	"time"

	"github.com/moul/go-clean-architecture/business-rules/responders/sprint"
)

type AddAssembler struct{}

func NewAddAssembler() *AddAssembler { return &AddAssembler{} }

func (a *AddAssembler) WriteAdd(resp sprintresp.AddSprintResponse) Add {
	add := Add{}
	add.CreatedAt = resp.GetCreatedAt()
	add.ID = resp.GetID()
	return add
}

type Add struct {
	CreatedAt time.Time `json:"created-at"`
	ID        int       `json:"id"`
}

func (g *Add) Serialize() []byte {
	out, _ := json.Marshal(*g)
	return out
}
