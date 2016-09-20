package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/moul/cleanarch/example/app/controllers/api"
	"github.com/moul/cleanarch/example/app/repos/sprints/gorm"

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
	controller := apicontrollers.New()
	controller.SetSprintsGateway(sprintsgorm.New(db))

	gin.GET("/sprints/:sprint-id", controller.GetSprint)
	gin.POST("/sprints", controller.AddSprint)

	// Start
	gin.Run()
}
