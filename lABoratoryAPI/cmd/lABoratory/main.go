package main

import (
	"lABoratory/lABoratoryAPI/internal/handlers"
	"lABoratory/lABoratoryAPI/internal/middleware"
	"lABoratory/lABoratoryAPI/internal/persistence/database"
	"lABoratory/lABoratoryAPI/internal/services"
	"lABoratory/lABoratoryAPI/internal/utils"
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

	router.GET("/assignment/:experimenttoken/:assignmentkey", assignmentHandler.GetAssignment)

	router.Use(middleware.ValidateJWT)

	router.POST("assignment/:experimenttoken/:assignmentkey", assignmentHandler.SetAssignment)

	router.GET("/assignments/overrides/:experimenttoken", assignmentHandler.GetOverrides)
	router.DELETE("assignment/:experimenttoken/:assignmentkey", assignmentHandler.DeleteAssignment)

	router.GET("assignments/:id", assignmentHandler.GetAssignmentsOfExperiment)

	router.GET("/user", authHandler.GetUser)
	router.DELETE("/user", authHandler.DeleteUser)

	router.GET("/experiments", experimentHandler.GetExperiments)
	router.POST("/experiments", experimentHandler.CreateExperiment)
	router.GET("/experiments/:id", experimentHandler.GetExperimentById)
	router.PUT("/experiments/:id", experimentHandler.UpdateExperiment)
	router.DELETE("/experiments/:id", experimentHandler.DeleteExperiment)

	url := "localhost:" + strconv.Itoa(utils.GetConfig().Port)
	log.Println("Listening on port " + strconv.Itoa(utils.GetConfig().Port))
	router.Run(url)
}
