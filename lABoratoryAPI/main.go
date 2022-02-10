package main

import (
	"lABoratory/lABoratoryAPI/config"
	"lABoratory/lABoratoryAPI/handlers"
	"log"
	"strconv"

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

	experimentHandler := handlers.NewExperimentHandler()

	// For release uncomment the next line
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/experiments", experimentHandler.GetExperiments)
	router.GET("/experiments/:id", experimentHandler.GetExperimentById)
	router.POST("/experiments", experimentHandler.CreateExperiment)
	router.PUT("/experiments/:id", experimentHandler.UpdateExperiment)
	router.DELETE("/experiments/:id", experimentHandler.DeleteExperiment)

	config, err := config.ReadConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	var url string = "localhost:" + strconv.Itoa(config.Port)

	router.Run(url)
}
