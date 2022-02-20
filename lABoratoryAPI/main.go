package main

import (
	"fmt"
	"lABoratory/lABoratoryAPI/config"
	"lABoratory/lABoratoryAPI/handlers"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies([]string{"0.0.0.0"})

	authHandler := handlers.NewAuthHandler()
	experimentHandler := handlers.NewExperimentHandler()

	router.GET("/")
	router.POST("/auth", authHandler.Authenticate)
	router.POST("/signup", authHandler.Signup)
	router.GET("/users", authHandler.GetUsers) // Only for debug

	//router.Use(middleware.ValidateJWT)

	router.GET("/user", authHandler.GetUser)
	router.DELETE("/user", authHandler.DeleteUser)

	router.GET("/experiments", experimentHandler.GetExperiments)
	router.GET("/experiments/:id", experimentHandler.GetExperimentById)
	router.POST("/experiments", experimentHandler.CreateExperiment)
	router.PUT("/experiments/:id", experimentHandler.UpdateExperiment)
	router.DELETE("/experiments/:id", experimentHandler.DeleteExperiment)

	url := "localhost:" + strconv.Itoa(config.ConfigParams.Port)
	fmt.Println("Listening on port " + strconv.Itoa(config.ConfigParams.Port))
	router.Run(url)
}
