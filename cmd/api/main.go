package main

import (
	"github.com/gin-gonic/gin"
	"github.com/moul/go-clean-architecture/app/controllers/api/sprint"
	"github.com/moul/go-clean-architecture/app/repositories/sprint/sprintrepomem"
	"github.com/moul/go-clean-architecture/business-rules/usecases/sprint/dto"
)

func main() {
	gin := gin.Default()

	controller := app_ctrl_api_sprint.NewController()
	controller.SetSprintGateway(sprintrepomem.New())
	controller.SetGetSprintResponseAssembler(sprintucdto.GetSprintResponseAssemblerDTO{})
	controller.SetGetSprintRequestBuilder(sprintucdto.GetSprintRequestBuilderDTO{})

	gin.GET("/sprints/:sprint-id", controller.GetSprint)
	gin.Run()
}
