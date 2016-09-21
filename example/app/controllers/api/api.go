package apicontrollers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/moul/cleanarch/example/bizrules/gateways"
	"github.com/moul/cleanarch/example/bizrules/usecases/add_sprint"
	"github.com/moul/cleanarch/example/bizrules/usecases/add_sprint/dto"
	"github.com/moul/cleanarch/example/bizrules/usecases/add_sprint/io"
	"github.com/moul/cleanarch/example/bizrules/usecases/get_sprint"
	"github.com/moul/cleanarch/example/bizrules/usecases/get_sprint/dto"
	"github.com/moul/cleanarch/example/bizrules/usecases/get_sprint/io"
)

type Controller struct {
	sprintsGw gateways.Sprints
}

func New() *Controller { return &Controller{} }

func (ctrl *Controller) SetSprintsGateway(val gateways.Sprints) { ctrl.sprintsGw = val }

func (ctrl *Controller) GetSprint(c *gin.Context) {
	sprintID, err := strconv.Atoi(c.Param("sprint-id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid 'sprint-id'"})
		return
	}

	respAsm := getsprintdto.ResponseAssembler{}
	uc := getsprint.New(ctrl.sprintsGw, respAsm)

	req := getsprintdto.RequestBuilder{}.
		Create().
		WithSprintID(sprintID).
		Build()

	resp, err := uc.Execute(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": (&GetAssembler{}).WriteGet(resp)})
}

func (ctrl *Controller) AddSprint(c *gin.Context) {
	respAsm := addsprintdto.ResponseAssembler{}
	uc := addsprint.New(ctrl.sprintsGw, respAsm)

	req := addsprintdto.RequestBuilder{}.
		Create().
		Build()

	resp, err := uc.Execute(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": (&AddAssembler{}).WriteAdd(resp)})
}

type AddAssembler struct{}

func (a *AddAssembler) WriteAdd(resp addsprintio.Response) Add {
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

type GetAssembler struct{}

func (a *GetAssembler) WriteGet(resp getsprintio.Response) Get {
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
