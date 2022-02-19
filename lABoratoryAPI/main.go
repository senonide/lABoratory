package main

import (
	"fmt"
	"lABoratory/lABoratoryAPI/config"
	"lABoratory/lABoratoryAPI/handlers"
	"lABoratory/lABoratoryAPI/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies([]string{"0.0.0.0"})

	authHandler := handlers.NewAuthHandler()
	experimentHandler := handlers.NewExperimentHandler()

	//Auth
	router.POST("/auth", authHandler.Authenticate)
	router.POST("/signup", authHandler.Singup)

	router.Use(middleware.ValidateJWT)

	// Experiments
	router.GET("/experiments", experimentHandler.GetExperiments)
	router.GET("/experiments/:id", experimentHandler.GetExperimentById)
	router.POST("/experiments", experimentHandler.CreateExperiment)
	router.PUT("/experiments/:id", experimentHandler.UpdateExperiment)
	router.DELETE("/experiments/:id", experimentHandler.DeleteExperiment)

	url := "localhost:" + strconv.Itoa(config.ConfigParams.Port)
	fmt.Println("Listening on port " + strconv.Itoa(config.ConfigParams.Port))
	router.Run(url)
}
