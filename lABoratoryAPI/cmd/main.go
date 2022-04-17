package main

import (
	"lABoratory/lABoratoryAPI/config"
	"lABoratory/lABoratoryAPI/handlers"
	"lABoratory/lABoratoryAPI/middleware"
	"lABoratory/lABoratoryAPI/persistence/database"
	"lABoratory/lABoratoryAPI/services"
	"lABoratory/lABoratoryAPI/utils"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies([]string{"0.0.0.0"})

	userRepository := database.NewDbUserRepository()
	experimentRepository := database.NewDbExperimentRepository()
	customerRepository := database.NewDbCustomerRepository()
	securityProvider := utils.NewSecurityProvider()

	assignmentService := services.NewAssignmentService(experimentRepository, customerRepository, securityProvider)
	experimentService := services.NewExperimentService(experimentRepository, securityProvider, assignmentService)
	authService := services.NewAuthService(userRepository, securityProvider, experimentService)

	authHandler := handlers.NewAuthHandler(authService)
	experimentHandler := handlers.NewExperimentHandler(experimentService, authService)
	assignmentHandler := handlers.NewAssignmentHandler(assignmentService)

	router.Use(middleware.CORSMiddleware)

	router.POST("/auth", authHandler.Authenticate)
	router.POST("/signup", authHandler.Signup)

	router.GET("/assignment/:key", assignmentHandler.GetAssignment)

	router.Use(middleware.ValidateJWT)

	router.PUT("assignment/:key", assignmentHandler.SetAssignment)

	router.GET("assignments/:id", assignmentHandler.GetAssignmentsOfExperiment)

	router.GET("/user", authHandler.GetUser)
	router.DELETE("/user", authHandler.DeleteUser)

	router.GET("/experiments", experimentHandler.GetExperiments)
	router.POST("/experiments", experimentHandler.CreateExperiment)
	router.GET("/experiments/:id", experimentHandler.GetExperimentById)
	router.PUT("/experiments/:id", experimentHandler.UpdateExperiment)
	router.DELETE("/experiments/:id", experimentHandler.DeleteExperiment)

	url := "localhost:" + strconv.Itoa(config.GetConfig().Port)
	log.Println("Listening on port " + strconv.Itoa(config.GetConfig().Port))
	router.Run(url)
}
