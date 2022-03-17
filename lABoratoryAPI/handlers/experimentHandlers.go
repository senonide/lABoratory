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

func NewExperimentHandler(es *services.ExperimentService, as *services.AuthService) *ExperimentHandler {
	eh := new(ExperimentHandler)
	eh.experimentService = es
	eh.authService = as
	return eh
}

func (eh *ExperimentHandler) GetExperiments(c *gin.Context) {
	tokenFromHeader := c.Request.Header.Get("Authorization")
	owner, err := eh.authService.GetOne(tokenFromHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	experiments, err := eh.experimentService.GetAll(owner)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, apitypes.GetExperimentsApiType(experiments))
}

func (eh *ExperimentHandler) GetExperimentById(c *gin.Context) {
	tokenFromHeader := c.Request.Header.Get("Authorization")
	owner, err := eh.authService.GetOne(tokenFromHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	id := c.Param("id")
	experiment, err := eh.experimentService.GetOne(id, owner)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	if experiment == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, apitypes.GetExperimentApiType(*experiment))
}

func (eh *ExperimentHandler) CreateExperiment(c *gin.Context) {
	var data models.Experiment
	err := c.BindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	tokenFromHeader := c.Request.Header.Get("Authorization")
	owner, err := eh.authService.GetOne(tokenFromHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	data.Owner = *owner
	err = eh.experimentService.Create(data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, apitypes.GetExperimentApiType(data))
}

func (eh *ExperimentHandler) UpdateExperiment(c *gin.Context) {
	tokenFromHeader := c.Request.Header.Get("Authorization")
	owner, err := eh.authService.GetOne(tokenFromHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	var data models.Experiment
	err = c.BindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	id := c.Param("id")
	data.Id = id
	err = eh.experimentService.Update(data, owner)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, apitypes.GetExperimentApiType(data))
}

func (eh *ExperimentHandler) DeleteExperiment(c *gin.Context) {
	tokenFromHeader := c.Request.Header.Get("Authorization")
	owner, err := eh.authService.GetOne(tokenFromHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	id := c.Param("id")
	wasDeleted, err := eh.experimentService.Delete(id, owner)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	if !wasDeleted {
		c.AbortWithStatus(http.StatusNotModified)
		return
	}
	c.IndentedJSON(http.StatusOK, responses.DeleteResponse{WasDeleted: wasDeleted})
}
