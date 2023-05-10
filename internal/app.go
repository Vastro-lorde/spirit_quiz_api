package internal

import (
	"fmt"
	"log"
	"net/http"
	"spirit_quiz/config"
	"spirit_quiz/internal/data/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	config, err := config.GetConfigs()
	if err != nil {
		fmt.Println(err)
	}
	router := gin.Default()

	// Add CORS middleware
	corConfig := cors.DefaultConfig()
	corConfig.AllowWildcard = true
	corConfig.AllowAllOrigins = true
	corConfig.AllowMethods = []string{"*"}
	corConfig.AllowCredentials = true
	corConfig.AllowHeaders = []string{"*"}

	router.Use(cors.New(corConfig))

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

	port := config.PORT
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
