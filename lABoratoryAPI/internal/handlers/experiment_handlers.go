package handlers

import (
	"lABoratory/lABoratoryAPI/internal/handlers/apitypes"
	"lABoratory/lABoratoryAPI/internal/handlers/responses"
	"lABoratory/lABoratoryAPI/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExperimentHandler struct {
	experimentService services.ExperimentServiceI
	authService       services.AuthServiceI
}

func NewExperimentHandler(es services.ExperimentServiceI, as services.AuthServiceI) *ExperimentHandler {
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
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
	var data apitypes.Experiment
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
	expModel := data.GetExperimentModel()
	expModel.Owner = *owner
	err = eh.experimentService.Create(expModel)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

func (eh *ExperimentHandler) UpdateExperiment(c *gin.Context) {
	tokenFromHeader := c.Request.Header.Get("Authorization")
	owner, err := eh.authService.GetOne(tokenFromHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	var data apitypes.Experiment
	err = c.BindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	id := c.Param("id")
	data.Id = id
	expModel := data.GetExperimentModel()
	err = eh.experimentService.Update(expModel, owner)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, data)
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.ResponseWithError{Message: "error", Error: err.Error()})
		return
	}
	if !wasDeleted {
		c.AbortWithStatus(http.StatusNotModified)
		return
	}
	c.IndentedJSON(http.StatusOK, responses.DeleteResponse{WasDeleted: wasDeleted})
}
