package handlers

import (
	"fmt"
	"lABoratory/lABoratoryAPI/models"
	"lABoratory/lABoratoryAPI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExperimentHandler struct {
}

type ExperimentHandlerI interface {
	GetExperiments(*gin.Context)
	GetExperimentById(*gin.Context)
	CreateExperiment(*gin.Context)
	UpdateExperiment(*gin.Context)
	DeleteExperiment(*gin.Context)
}

func NewExperimentHandler() *ExperimentHandler {
	var eh *ExperimentHandler

	return eh
}

func (eh *ExperimentHandler) GetExperiments(c *gin.Context) {

	experiments, err := services.Read()
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, experiments)
}

func (eh *ExperimentHandler) GetExperimentById(c *gin.Context) {
	id := c.Param("id")
	experiment, err := services.ReadOne(id)

	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, experiment)
}

func (eh *ExperimentHandler) CreateExperiment(c *gin.Context) {
	var newExperiment models.Experiment

	err := c.BindJSON(&newExperiment)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	services.Create(newExperiment)

	c.IndentedJSON(http.StatusCreated, newExperiment)
}

func (eh *ExperimentHandler) UpdateExperiment(c *gin.Context) {

	var newExperiment models.Experiment
	id := c.Param("id")

	err := c.BindJSON(&newExperiment)
	if err != nil {
		return
	}

	services.Update(newExperiment, id)

	c.IndentedJSON(http.StatusOK, newExperiment)

}

func (eh *ExperimentHandler) DeleteExperiment(c *gin.Context) {

	id := c.Param("id")
	services.Delete(id)

	experiments, err := services.Read()
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, experiments)

}
