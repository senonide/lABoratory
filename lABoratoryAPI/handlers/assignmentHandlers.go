package handlers

import (
	"lABoratory/lABoratoryAPI/handlers/responses"
	"lABoratory/lABoratoryAPI/services"
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
	assignment, err := ah.service.GetAssignment(c.Param("key"))
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
	err = ah.service.SetAssignment(c.Param("key"), assignmentName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, assignmentName)
}
