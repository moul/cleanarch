package apicontrollers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/moul/cleanarch/example/bizrules/usecases/add_sprint"
	"github.com/moul/cleanarch/example/bizrules/usecases/add_sprint/dto"
)

type AddSprint struct {
	uc *addsprint.UseCase
}

func NewAddSprint(uc *addsprint.UseCase) *AddSprint { return &AddSprint{uc: uc} }

func (ctrl *AddSprint) Execute(ctx *gin.Context) {
	req := addsprintdto.RequestBuilder{}.
		Create().
		Build()

	resp, err := ctrl.uc.Execute(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": AddSprintResponse{
		CreatedAt: resp.GetCreatedAt(),
		ID:        resp.GetID(),
	}})
}

type AddSprintResponse struct {
	CreatedAt time.Time `json:"created-at"`
	ID        int       `json:"id"`
}
