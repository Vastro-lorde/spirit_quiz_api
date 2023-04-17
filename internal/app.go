package internal

import (
	"fmt"
	"log"
	"net/http"
	"spirit_quiz/config"
	"spirit_quiz/internal/data/database"
	"spirit_quiz/internal/handlers"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	config, err := config.GetConfigs()
	if err != nil {
		fmt.Println(err)
	}
	router := gin.Default()

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

	router.POST("/create/category", handlers.CreateCategory)
	router.POST("/create/question", handlers.CreateQuestion)

	router.GET("/get/categories", handlers.GetCategories)
	router.GET("/get/questions/:id", handlers.GetQuestionsByCategoryId)

	port := config.PORT
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
