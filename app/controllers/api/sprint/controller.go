package app_ctrl_api_sprint

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moul/go-clean-architecture/app/view-models/api/sprint"
	"github.com/moul/go-clean-architecture/business-rules/gateways/sprint"
	"github.com/moul/go-clean-architecture/business-rules/requestors/sprint"
	"github.com/moul/go-clean-architecture/business-rules/responders/sprint"
	"github.com/moul/go-clean-architecture/business-rules/usecases/sprint"
)

type Controller struct {
	sprintGateway              sprintgw.SprintGateway
	getSprintResponseAssembler sprintresp.GetSprintResponseAssembler
	getSprintRequestBuilder    sprintreq.GetSprintRequestBuilder
	addSprintResponseAssembler sprintresp.AddSprintResponseAssembler
	addSprintRequestBuilder    sprintreq.AddSprintRequestBuilder

	getViewModelAssembler app_vm_api_sprint.GetAssembler
	addViewModelAssembler app_vm_api_sprint.AddAssembler
}

func NewController() *Controller {
	return &Controller{}
}

// controller boundaries
func (ctrl *Controller) SetSprintGateway(val sprintgw.SprintGateway) { ctrl.sprintGateway = val }
func (ctrl *Controller) SetGetSprintResponseAssembler(val sprintresp.GetSprintResponseAssembler) {
	ctrl.getSprintResponseAssembler = val
}
func (ctrl *Controller) SetGetSprintRequestBuilder(val sprintreq.GetSprintRequestBuilder) {
	ctrl.getSprintRequestBuilder = val
}
func (ctrl *Controller) SetAddSprintResponseAssembler(val sprintresp.AddSprintResponseAssembler) {
	ctrl.addSprintResponseAssembler = val
}
func (ctrl *Controller) SetAddSprintRequestBuilder(val sprintreq.AddSprintRequestBuilder) {
	ctrl.addSprintRequestBuilder = val
}

// GetSprint controller
func (ctrl *Controller) GetSprint(c *gin.Context) {
	sprintID, err := strconv.Atoi(c.Param("sprint-id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid 'sprint-id'"})
		return
	}

	uc := sprintuc.NewGetSprint()
	uc.SetSprintGateway(ctrl.sprintGateway)
	uc.SetGetSprintResponseAssembler(ctrl.getSprintResponseAssembler)

	req := ctrl.getSprintRequestBuilder.Create().WithSprintID(sprintID).Build()

	resp, err := uc.Execute(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": ctrl.getViewModelAssembler.WriteGet(resp)})
}

// AddSPrint controller
func (ctrl *Controller) AddSprint(c *gin.Context) {
	uc := sprintuc.NewAddSprint()
	uc.SetSprintGateway(ctrl.sprintGateway)
	uc.SetGetSprintResponseAssembler(ctrl.addSprintResponseAssembler)

	req := ctrl.addSprintRequestBuilder.Create().Build()

	resp, err := uc.Execute(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": ctrl.addViewModelAssembler.WriteAdd(resp)})
}
