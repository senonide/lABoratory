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

type IExperimentHandler interface {
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
	experiments, err := eh.service.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusOK, experiments)
}

func (eh *ExperimentHandler) GetExperimentById(c *gin.Context) {
	id := c.Param("id")
	experiment, err := eh.service.GetOne(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if experiment == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, experiment)
}

func (eh *ExperimentHandler) CreateExperiment(c *gin.Context) {
	var newExperiment models.Experiment
	err := c.BindJSON(&newExperiment)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	errCreating := eh.service.Create(newExperiment)
	if errCreating != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusCreated, newExperiment)
}

func (eh *ExperimentHandler) UpdateExperiment(c *gin.Context) {
	var newExperiment models.Experiment
	id := c.Param("id")
	err := c.BindJSON(&newExperiment)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	errUpdating := eh.service.Update(newExperiment, id)
	if errUpdating != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusOK, newExperiment)
}

func (eh *ExperimentHandler) DeleteExperiment(c *gin.Context) {
	id := c.Param("id")
	wasDeleted, errDeleting := eh.service.Delete(id)
	if errDeleting != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if !wasDeleted {
		c.AbortWithStatus(http.StatusNotModified)
		return
	}
	c.IndentedJSON(http.StatusOK, wasDeleted)
}
