package handlers

import (
	"net/http"
	dtos "spirit_quiz/internal/Dtos"
	"spirit_quiz/internal/data/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateResult(context *gin.Context) {

	var question dtos.CreateQuestionDto
	if err := context.ShouldBindJSON(&question); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var newQuestion = models.Question{
		ID:         uuid.New(),
		Text:       question.Text,
		CategoryID: uuid.MustParse(question.CategoryID),
	}
	if err := db.Create(&newQuestion).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusCreated, newQuestion)
}
