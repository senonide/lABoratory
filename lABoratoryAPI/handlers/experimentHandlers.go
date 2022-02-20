package handlers

import (
	"lABoratory/lABoratoryAPI/handlers/apitypes"
	"lABoratory/lABoratoryAPI/handlers/responses"
	"lABoratory/lABoratoryAPI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExperimentHandler struct {
	service *services.ExperimentService
}

func NewExperimentHandler() *ExperimentHandler {
	eh := new(ExperimentHandler)
	eh.service = services.NewExperimentService()
	return eh
}

func (eh *ExperimentHandler) GetExperiments(c *gin.Context) {
	experiments, err := eh.service.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, experiments)
}

func (eh *ExperimentHandler) GetExperimentById(c *gin.Context) {
	id := c.Param("id")
	experiment, err := eh.service.GetOne(id)
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
	var newExperiment apitypes.Experiment
	err := c.BindJSON(&newExperiment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	errCreating := eh.service.Create(newExperiment.GetExperimentModel())
	if errCreating != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: errCreating.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newExperiment)
}

func (eh *ExperimentHandler) UpdateExperiment(c *gin.Context) {
	var newExperiment apitypes.Experiment
	err := c.BindJSON(&newExperiment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	errUpdating := eh.service.Update(newExperiment.GetExperimentModel())
	if errUpdating != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: errUpdating.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, newExperiment)
}

func (eh *ExperimentHandler) DeleteExperiment(c *gin.Context) {
	id := c.Param("id")
	wasDeleted, errDeleting := eh.service.Delete(id)
	if errDeleting != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: errDeleting.Error()})
		return
	}
	if !wasDeleted {
		c.AbortWithStatus(http.StatusNotModified)
		return
	}
	c.IndentedJSON(http.StatusOK, responses.DeleteResponse{WasDeleted: wasDeleted})
}
