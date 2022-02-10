package handlers

import (
	"lABoratory/lABoratoryAPI/models"
	"lABoratory/lABoratoryAPI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExperimentHandler struct {
	service *services.ExperimentService
}

type ExperimentHandlerI interface {
	GetExperiments(*gin.Context, *services.ExperimentService)
	GetExperimentById(*gin.Context, *services.ExperimentService)
	CreateExperiment(*gin.Context, *services.ExperimentService)
	UpdateExperiment(*gin.Context, *services.ExperimentService)
	DeleteExperiment(*gin.Context, *services.ExperimentService)
}

func NewExperimentHandler() *ExperimentHandler {
	eh := new(ExperimentHandler)
	eh.service = services.NewExperimentService()
	return eh
}

func (eh *ExperimentHandler) GetExperiments(c *gin.Context) {

	experiments, err := eh.service.Read()
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, experiments)
}

func (eh *ExperimentHandler) GetExperimentById(c *gin.Context) {
	id := c.Param("id")
	experiment, err := eh.service.ReadOne(id)

	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, experiment)
}

func (eh *ExperimentHandler) CreateExperiment(c *gin.Context) {
	var newExperiment models.Experiment

	err := c.BindJSON(&newExperiment)
	if err != nil {
		return
	}

	eh.service.Create(newExperiment)

	c.IndentedJSON(http.StatusCreated, newExperiment)
}

func (eh *ExperimentHandler) UpdateExperiment(c *gin.Context) {

	var newExperiment models.Experiment
	id := c.Param("id")

	err := c.BindJSON(&newExperiment)
	if err != nil {
		return
	}

	eh.service.Update(newExperiment, id)

	c.IndentedJSON(http.StatusOK, newExperiment)

}

func (eh *ExperimentHandler) DeleteExperiment(c *gin.Context) {

	id := c.Param("id")
	eh.service.Delete(id)

	experiments, err := eh.service.Read()
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, experiments)

}
