package internal

import (
	"fmt"
	"log"
	"net/http"
	"spirit_quiz/config"
	"spirit_quiz/internal/data/database"
	"spirit_quiz/internal/services"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	config, err := config.GetConfigs()
	if err != nil {
		fmt.Println(err)
	}
	router := gin.Default()
	router.Use(services.CorsMiddleware()) // adding cors to the pipeline

	// router.Use(cors.New(cors.Config{
	// 	AllowWildcard:    true,
	// 	AllowOrigins:     []string{"http://localhost:3000", "*"},
	// 	AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
	// 	AllowHeaders:     []string{"*"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	//Connect to database
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Close database connection when main function exits
	defer func() {
		dbConn, err := db.DB()
		if err != nil {
			log.Println("error getting database connection:", err)
			return
		}
		if err := dbConn.Close(); err != nil {
			log.Println("error closing database connection:", err)
		}
	}()

	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "welcome to quiz api")
	})

	authRoutes := router.Group("/auth")
	SetupAuthRoutes(authRoutes)

	quizRoutes := router.Group("/quiz")
	SetupQuizRoutes(quizRoutes)

	imageRoutes := router.Group("/image")
	SetupImageRoutes(imageRoutes)

	userRoutes := router.Group("/user")
	SetupUserRoutes(userRoutes)

	resultRoutes := router.Group("/result")
	SetupResultRoutes(resultRoutes)

	fmt.Println("before routes")
	port := config.PORT
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
