package handlers

import (
	"lABoratory/lABoratoryAPI/internal/handlers/responses"
	"lABoratory/lABoratoryAPI/internal/models"
	"lABoratory/lABoratoryAPI/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AssignmentHandler struct {
	service services.AssignmentServiceI
}

func NewAssignmentHandler(as services.AssignmentServiceI) *AssignmentHandler {
	ah := new(AssignmentHandler)
	ah.service = as
	return ah
}

func (ah *AssignmentHandler) GetAssignment(c *gin.Context) {
	experimentToken := c.Param("experimenttoken")
	assignmentKey := c.Param("assignmentkey")
	assignment, err := ah.service.GetAssignment(experimentToken, assignmentKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, assignment)
}

func (ah *AssignmentHandler) SetAssignment(c *gin.Context) {
	var assignmentName string
	err := c.BindJSON(&assignmentName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	err = ah.service.SetAssignment(c.Param("assignmentkey"), c.Param("experimenttoken"), assignmentName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, assignmentName)
}

func (ah *AssignmentHandler) GetAssignmentsOfExperiment(c *gin.Context) {
	assignments, err := ah.service.GetAssignmentsOfExperiment(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, assignments)
}

func (ah *AssignmentHandler) GetOverrides(c *gin.Context) {
	experimentToken := c.Param("experimenttoken")
	assignments, err := ah.service.GetAssignmentsOfExperiment(experimentToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	overrides := []models.Customer{}
	for _, a := range assignments {
		if a.Override {
			overrides = append(overrides, a)
		}
	}
	c.IndentedJSON(http.StatusOK, overrides)
}

func (ah *AssignmentHandler) DeleteAssignment(c *gin.Context) {
	experimentToken := c.Param("experimenttoken")
	assignmentKey := c.Param("assignmentkey")
	wasDeleted, err := ah.service.DeleteAssignment(experimentToken, assignmentKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, wasDeleted)
}
