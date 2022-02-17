package main

import (
	"fmt"
	"lABoratory/lABoratoryAPI/config"
	"lABoratory/lABoratoryAPI/handlers"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	experimentHandler := handlers.NewExperimentHandler()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.SetTrustedProxies([]string{"0.0.0.0"})
	router.GET("/experiments", experimentHandler.GetExperiments)
	router.GET("/experiments/:id", experimentHandler.GetExperimentById)
	router.POST("/experiments", experimentHandler.CreateExperiment)
	router.PUT("/experiments/:id", experimentHandler.UpdateExperiment)
	router.DELETE("/experiments/:id", experimentHandler.DeleteExperiment)

	url := "localhost:" + strconv.Itoa(config.ConfigParams.Port)
	fmt.Println("Listening on port " + strconv.Itoa(config.ConfigParams.Port))
	router.Run(url)
}
