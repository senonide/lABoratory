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
	assignment, err := ah.service.GetAssignment(c.Param("token"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, assignment)
}
