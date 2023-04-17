package handlers

import (
	"net/http"
	dtos "spirit_quiz/internal/Dtos"
	"spirit_quiz/internal/data/database"
	"spirit_quiz/internal/data/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

var db, err = database.Connect()

func CreateCategory(context *gin.Context) {

	var category dtos.CreateCategoryDto
	if err := context.ShouldBindJSON(&category); err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var newCategory = models.Category{
		ID: uuid.New(),
		Name: category.Name,
	}
	if err := db.Create(&newCategory).Error; err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	context.JSON(http.StatusCreated, newCategory)
}

func CreateQuestion(context *gin.Context) {

	var question dtos.CreateQuestionDto
	if err := context.ShouldBindJSON(&question); err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var newQuestion = models.Question{
		ID: uuid.New(),
		Text: question.Text,
		CategoryID: uuid.MustParse(question.CategoryID),
	}
	if err := db.Create(&newQuestion).Error; err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	context.JSON(http.StatusCreated, newQuestion)
}

func GetCategories(context *gin.Context) {
	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var responseDto []dtos.CategoryResponseDto

	for _, category := range categories {
		var categoryDto dtos.CategoryResponseDto
		err := mapstructure.Decode(category, &categoryDto)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		responseDto = append(responseDto, categoryDto)
	}
	context.JSON(http.StatusOK, responseDto)
}

func GetQuestionsByCategoryId(context *gin.Context) {
	var questions []models.Question
	if err := db.Where("category_id = ?", context.Params.ByName("id")).Find(&questions).Error; err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var responseDto []dtos.QuestionResponseDto

	for _, question := range questions {
		var questionResponseDto dtos.QuestionResponseDto
		err := mapstructure.Decode(question, &questionResponseDto)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}
		responseDto = append(responseDto, questionResponseDto)
	}
	context.JSON(http.StatusOK, responseDto)
}
