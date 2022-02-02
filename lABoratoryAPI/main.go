package main

import (
	"net/http"

	"lABoratory/lABoratoryAPI/models"
	"lABoratory/lABoratoryAPI/services"

	"github.com/gin-gonic/gin"
)

// Example data
/*
var prueba = []models.Experiment{
	{Id: "1", Name: "Experiment 1", C: 50, ActiveExperiments: []int{25, 25}},
	{Id: "2", Name: "Experiment 2", C: 50, ActiveExperiments: []int{50}},
	{Id: "3", Name: "Experiment 3", C: 50, ActiveExperiments: []int{20, 30}},
	{Id: "4", Name: "Experiment 4", C: 50, ActiveExperiments: []int{10, 20, 20}},
	{Id: "5", Name: "Experiment 5", C: 50, ActiveExperiments: []int{25, 25}},
	{Id: "6", Name: "Experiment 6", C: 70, ActiveExperiments: []int{25, 0}},
}*/

func main() {
	// For release uncomment the next line
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/experiments", getExperiments)
	router.GET("/experiments/:id", getExperimentById)
	router.POST("/experiments", createExperiment)
	router.PUT("/experiments/:id", updateExperiment)
	router.DELETE("/experiments/:id", deleteExperiment)

	router.Run("localhost:8080")
}

func getExperiments(c *gin.Context) {

	experiments, err := services.Read()
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, experiments)
}

func getExperimentById(c *gin.Context) {
	id := c.Param("id")
	experiment, err := services.ReadOne(id)

	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, experiment)
}

func createExperiment(c *gin.Context) {
	var newExperiment models.Experiment

	err := c.BindJSON(&newExperiment)
	if err != nil {
		return
	}

	services.Create(newExperiment)

	c.IndentedJSON(http.StatusCreated, newExperiment)
}

func updateExperiment(c *gin.Context) {

	var newExperiment models.Experiment
	id := c.Param("id")

	err := c.BindJSON(&newExperiment)
	if err != nil {
		return
	}

	services.Update(newExperiment, id)

	c.IndentedJSON(http.StatusOK, newExperiment)

}

func deleteExperiment(c *gin.Context) {

	id := c.Param("id")
	services.Delete(id)

	experiments, err := services.Read()
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, experiments)

}
