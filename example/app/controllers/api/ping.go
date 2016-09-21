package apicontrollers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moul/cleanarch/example/bizrules/usecases/ping"
	"github.com/moul/cleanarch/example/bizrules/usecases/ping/dto"
)

type Ping struct {
	uc *ping.UseCase
}

func NewPing(uc *ping.UseCase) *Ping {
	return &Ping{uc: uc}
}

func (ctrl *Ping) Execute(ctx *gin.Context) {
	req := pingdto.RequestBuilder{}.Create().Build()

	resp, err := ctrl.uc.Execute(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": PingResponse{
		Pong: resp.GetPong(),
	}})
}

type PingResponse struct {
	Pong string `json:"pong"`
}
