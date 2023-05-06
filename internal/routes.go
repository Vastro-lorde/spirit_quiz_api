package internal

import (
	"spirit_quiz/internal/handlers"

	"github.com/gin-gonic/gin"
)

// Auth Routes
func SetupAuthRoutes(router *gin.RouterGroup) {
	router.POST("/register", handlers.SignUp)
	router.POST("/verify-email", handlers.VerifyEmail)
	router.POST("/login", handlers.Login)
	router.POST("/request-change-password", handlers.RequestResetPassword)
	router.POST("/reset-password", handlers.ResetPassword)
}

func SetupQuizRoutes(router *gin.RouterGroup) {

	//create routes
	createRoutes := router.Group("/create")
	{
		createRoutes.POST("/category", handlers.CreateCategory)
		createRoutes.POST("/question", handlers.CreateQuestion)
		createRoutes.POST("/option", handlers.CreateOption)
	}

	router.GET("categories", handlers.GetCategories)

	router.GET("/questions", handlers.GetQuestions)
	router.GET("/questions/:id", handlers.GetQuestionsByCategoryId)

	router.GET("/options", handlers.GetOptions)
	router.GET("/options/:id", handlers.GetOptionsByQuestionId)

}
