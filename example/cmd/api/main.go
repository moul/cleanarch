package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/moul/cleanarch/example/app/controllers/api"
	"github.com/moul/cleanarch/example/app/repos/sprints/gorm"
	"github.com/moul/cleanarch/example/app/repos/sprints/mem"
	"github.com/moul/cleanarch/example/bizrules/gateways"
	"github.com/moul/cleanarch/example/bizrules/usecases/add_sprint"
	"github.com/moul/cleanarch/example/bizrules/usecases/add_sprint/dto"
	"github.com/moul/cleanarch/example/bizrules/usecases/get_sprint"
	"github.com/moul/cleanarch/example/bizrules/usecases/get_sprint/dto"
	"github.com/moul/cleanarch/example/bizrules/usecases/ping"
	"github.com/moul/cleanarch/example/bizrules/usecases/ping/dto"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Setup gateways
	var sprintsGw gateways.Sprints
	if len(os.Args) > 1 && os.Args[1] == "--mem" {
		// configure a memory-based sprints gateway
		sprintsGw = sprintsmem.New()
	} else {
		// configure a sqlite-based sprints gateway
		db, err := gorm.Open("sqlite3", "test.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()
		sprintsGw = sprintsgorm.New(db)
	}

	// Setup usecases
	getSprint := getsprint.New(sprintsGw, getsprintdto.ResponseAssembler{})
	addSprint := addsprint.New(sprintsGw, addsprintdto.ResponseAssembler{})
	ping := ping.New(pingdto.ResponseAssembler{})
	//closeSprint := closesprint.New(sprintsGw, closesprintdto.ResponseBuilder{})

	// Setup API
	gin := gin.Default()
	gin.GET("/sprints/:sprint-id", apicontrollers.NewGetSprint(&getSprint).Execute)
	gin.POST("/sprints", apicontrollers.NewAddSprint(&addSprint).Execute)
	gin.GET("/ping", apicontrollers.NewPing(&ping).Execute)
	//gin.DELETE("/sprints/:sprint-id", apicontrollers.NewCloseSprint(&closeSprint).Execute)

	// Start
	gin.Run()
}
