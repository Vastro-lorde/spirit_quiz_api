package internal

import (
	"spirit_quiz/internal/handlers"
	"spirit_quiz/internal/middlewares"

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
	// Apply the Auth middleware to all routes in the router group
	router.Use(middlewares.Auth())

	//create routes
	creationRoutes := router.Group("/create")
	{
		creationRoutes.Use(middlewares.RoleAuth("ADMIN"))
		creationRoutes.POST("/category", handlers.CreateCategory)
		creationRoutes.POST("/question", handlers.CreateQuestion)
		creationRoutes.POST("/option", handlers.CreateOption)
	}

	categoryRoutes := router.Group("/categories")
	{
		categoryRoutes.GET("/", handlers.GetCategories)
		categoryRoutes.PATCH("/:id", middlewares.RoleAuth("ADMIN"), handlers.UpdateCategoryById)
		categoryRoutes.DELETE("/:id", middlewares.RoleAuth("ADMIN"), handlers.DeleteCategoryById)
	}

	questionRoutes := router.Group("/questions")
	{
		questionRoutes.GET("/", handlers.GetQuestions)
		questionRoutes.GET("/:id", handlers.GetQuestionsByCategoryId)
		questionRoutes.PATCH("/:id", middlewares.RoleAuth("ADMIN"), handlers.UpdateQuestionById)
		questionRoutes.DELETE("/:id", middlewares.RoleAuth("ADMIN"), handlers.DeleteQuestionById)
	}

	optionRoutes := router.Group("/options")
	{
		optionRoutes.GET("/", handlers.GetOptions)
		optionRoutes.GET("/:id", handlers.GetOptionsByQuestionId)
		optionRoutes.PATCH("/:id", middlewares.RoleAuth("ADMIN"), handlers.UpdateOptionById)
		optionRoutes.DELETE("/:id", middlewares.RoleAuth("ADMIN"), handlers.DeleteOptionById)
	}

}
