package main

import (
	"lABoratory/lABoratoryAPI/config"
	"lABoratory/lABoratoryAPI/handlers"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

	var url string = "localhost:" + strconv.Itoa(config.ConfigParams.Port)

	router.Run(url)
}
