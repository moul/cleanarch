package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/moul/go-clean-architecture/app/controllers/api/sprint"
	"github.com/moul/go-clean-architecture/app/repositories/sprint/sprintrepogorm"
	"github.com/moul/go-clean-architecture/business-rules/usecases/sprint/dto"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize database
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Setup API
	gin := gin.Default()
	controller := app_ctrl_api_sprint.NewController()
	controller.SetSprintGateway(app_repo_sprint_gorm.New(db))
	controller.SetGetSprintResponseAssembler(sprintucdto.GetSprintResponseAssemblerDTO{})
	controller.SetGetSprintRequestBuilder(sprintucdto.GetSprintRequestBuilderDTO{})
	controller.SetAddSprintResponseAssembler(sprintucdto.AddSprintResponseAssemblerDTO{})
	controller.SetAddSprintRequestBuilder(sprintucdto.AddSprintRequestBuilderDTO{})

	gin.GET("/sprints/:sprint-id", controller.GetSprint)
	gin.POST("/sprints", controller.AddSprint)

	// Start
	gin.Run()
}
