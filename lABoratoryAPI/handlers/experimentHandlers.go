package handlers

import (
	"lABoratory/lABoratoryAPI/handlers/apitypes"
	"lABoratory/lABoratoryAPI/handlers/responses"
	"lABoratory/lABoratoryAPI/models"
	"lABoratory/lABoratoryAPI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExperimentHandler struct {
	experimentService *services.ExperimentService
	authService       *services.AuthService
}

func NewExperimentHandler() *ExperimentHandler {
	eh := new(ExperimentHandler)
	eh.experimentService = services.NewExperimentService()
	eh.authService = services.NewAuthService()
	return eh
}

func (eh *ExperimentHandler) GetExperiments(c *gin.Context) {
	experiments, err := eh.experimentService.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, apitypes.GetExperimentsApiType(experiments))
}

func (eh *ExperimentHandler) GetExperimentById(c *gin.Context) {
	id := c.Param("id")
	experiment, err := eh.experimentService.GetOne(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	if experiment == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, apitypes.GetExperimentApiType(*experiment))
}

func (eh *ExperimentHandler) CreateExperiment(c *gin.Context) {
	var data map[string]interface{}
	err := c.BindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	assignmentsData := data["assignments"].([]interface{})
	assignments := []models.Assignment{}
	for _, assignment := range assignmentsData {
		assignments = append(assignments, models.Assignment{
			AssignmentName:  assignment.(map[string]interface{})["assignmentName"].(string),
			AssignmentValue: assignment.(map[string]interface{})["assignmentValue"].(float64),
		})
	}
	tokenFromCookie, err := c.Cookie("jwt")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	owner, err := eh.authService.GetOne(tokenFromCookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	newExperiment := models.Experiment{Name: data["name"].(string), Assignments: assignments, Owner: *owner}
	err = eh.experimentService.Create(newExperiment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, apitypes.GetExperimentApiType(newExperiment))
}

func (eh *ExperimentHandler) UpdateExperiment(c *gin.Context) {
	var data map[string]interface{}
	err := c.BindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	assignmentsData := data["assignments"].([]interface{})
	assignments := []models.Assignment{}
	for _, assignment := range assignmentsData {
		assignments = append(assignments, models.Assignment{
			AssignmentName:  assignment.(map[string]interface{})["assignmentName"].(string),
			AssignmentValue: assignment.(map[string]interface{})["assignmentValue"].(float64),
		})
	}
	id := c.Param("id")
	newExperiment := models.Experiment{Id: id, Name: data["name"].(string), Assignments: assignments}
	err = eh.experimentService.Update(newExperiment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, apitypes.GetExperimentApiType(newExperiment))
}

func (eh *ExperimentHandler) DeleteExperiment(c *gin.Context) {
	id := c.Param("id")
	wasDeleted, err := eh.experimentService.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	if !wasDeleted {
		c.AbortWithStatus(http.StatusNotModified)
		return
	}
	c.IndentedJSON(http.StatusOK, responses.DeleteResponse{WasDeleted: wasDeleted})
}
