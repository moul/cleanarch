package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/moul/cleanarch/example/app/controllers/api"
	"github.com/moul/cleanarch/example/app/repos/sprints/gorm"
	"github.com/moul/cleanarch/example/bizrules/usecases/add_sprint"
	"github.com/moul/cleanarch/example/bizrules/usecases/add_sprint/dto"
	"github.com/moul/cleanarch/example/bizrules/usecases/get_sprint"
	"github.com/moul/cleanarch/example/bizrules/usecases/get_sprint/dto"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize database
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Setup gateways
	sprintsGw := sprintsgorm.New(db)

	// Setup usecases
	getSprint := getsprint.New(sprintsGw, getsprintdto.ResponseAssembler{})
	addSprint := addsprint.New(sprintsGw, addsprintdto.ResponseAssembler{})
	//closeSprint := closesprint.New(sprintsGw, closesprintdto.ResponseBuilder{})

	// Setup API
	gin := gin.Default()
	gin.GET("/sprints/:sprint-id", apicontrollers.NewGetSprint(&getSprint).Execute)
	gin.POST("/sprints", apicontrollers.NewAddSprint(&addSprint).Execute)
	//gin.DELETE("/sprints/:sprint-id", apicontrollers.NewCloseSprint(&closeSprint).Execute)

	// Start
	gin.Run()
}
