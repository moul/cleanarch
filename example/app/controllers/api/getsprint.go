package apicontrollers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/moul/cleanarch/example/bizrules/usecases/get_sprint"
	"github.com/moul/cleanarch/example/bizrules/usecases/get_sprint/dto"
)

type GetSprint struct {
	uc *getsprint.UseCase
}

func NewGetSprint(uc *getsprint.UseCase) *GetSprint {
	return &GetSprint{uc: uc}
}

func (ctrl *GetSprint) Execute(ctx *gin.Context) {
	sprintID, err := strconv.Atoi(ctx.Param("sprint-id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid 'sprint-id'"})
		return
	}

	req := getsprintdto.RequestBuilder{}.
		Create().
		WithSprintID(sprintID).
		Build()

	resp, err := ctrl.uc.Execute(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": GetSprintResponse{
		CreatedAt:         resp.GetCreatedAt(),
		EffectiveClosedAt: resp.GetEffectiveClosedAt(),
		ExpectedClosedAt:  resp.GetExpectedClosedAt(),
		Status:            resp.GetStatus(),
	}})
}

type GetSprintResponse struct {
	CreatedAt         time.Time `json:"created-at"`
	EffectiveClosedAt time.Time `json:"effective-closed-at"`
	ExpectedClosedAt  time.Time `json:"expected-closed-at"`
	Status            string    `json:"status"`
}
