package main

import (
	"fmt"
	"lABoratory/lABoratoryAPI/config"
	"lABoratory/lABoratoryAPI/handlers"
	"lABoratory/lABoratoryAPI/middleware"
	"lABoratory/lABoratoryAPI/persistence/database"
	"lABoratory/lABoratoryAPI/services"
	"lABoratory/lABoratoryAPI/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies([]string{"0.0.0.0"})

	userRepository := database.NewDbUserRepository()
	experimetnRepository := database.NewDbExperimentRepository()
	securityProvider := utils.NewSecurityProvider()

	experimentService := services.NewExperimentService(experimetnRepository)
	authService := services.NewAuthService(userRepository, securityProvider, experimentService)

	authHandler := handlers.NewAuthHandler(authService)
	experimentHandler := handlers.NewExperimentHandler(experimentService, authService)

	router.Use(middleware.CORSMiddleware())

	router.POST("/auth", authHandler.Authenticate)
	router.POST("/signup", authHandler.Signup)
	router.GET("/users", authHandler.GetUsers) // Only for debug

	router.Use(middleware.ValidateJWT)

	router.GET("/user", authHandler.GetUser)
	router.DELETE("/user", authHandler.DeleteUser)

	router.GET("/experiments", experimentHandler.GetExperiments)
	router.POST("/experiments", experimentHandler.CreateExperiment)
	router.GET("/experiments/:id", experimentHandler.GetExperimentById)
	router.PUT("/experiments/:id", experimentHandler.UpdateExperiment)
	router.DELETE("/experiments/:id", experimentHandler.DeleteExperiment)

	url := "localhost:" + strconv.Itoa(config.ConfigParams.Port)
	fmt.Println("Listening on port " + strconv.Itoa(config.ConfigParams.Port))
	router.Run(url)
}
