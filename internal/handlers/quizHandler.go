package handlers

import (
	"math/rand"
	"net/http"
	dtos "spirit_quiz/internal/Dtos"
	"spirit_quiz/internal/data/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

// category handlers
func CreateCategory(context *gin.Context) {

	var category dtos.CreateCategoryDto
	if err := context.ShouldBindJSON(&category); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var newCategory = models.Category{
		ID:   uuid.New(),
		Name: category.Name,
	}
	if err := db.Create(&newCategory).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.AbortWithStatusJSON(http.StatusCreated, newCategory)
}

func GetCategories(context *gin.Context) {
	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	var responseDto []dtos.CategoryResponseDto

	for _, category := range categories {
		var categoryDto dtos.CategoryResponseDto
		err := mapstructure.Decode(category, &categoryDto)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}
		responseDto = append(responseDto, categoryDto)
	}
	context.AbortWithStatusJSON(http.StatusOK, responseDto)
}

// question handlers
func CreateQuestion(context *gin.Context) {

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

func GetQuestions(context *gin.Context) {
	var questions []models.Question
	if err := db.Find(&questions).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	var responseDto []dtos.QuestionResponseDto

	for _, question := range questions {
		var questionResponseDto dtos.QuestionResponseDto
		err := mapstructure.Decode(question, &questionResponseDto)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}
		responseDto = append(responseDto, questionResponseDto)
	}
	context.AbortWithStatusJSON(http.StatusOK, responseDto)
}

func GetQuestionsByCategoryId(context *gin.Context) {
	var questions []models.Question
	if err := db.Where("category_id = ?", context.Params.ByName("id")).Find(&questions).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Shuffle the questions
	for i := range questions {
		j := rand.Intn(i + 1)
		questions[i], questions[j] = questions[j], questions[i]
	}

	// Return up to 20 questions
	if len(questions) > 20 {
		questions = questions[:20]
	}

	var responseDto []dtos.QuestionResponseDto

	for _, question := range questions {
		var questionResponseDto dtos.QuestionResponseDto
		err := mapstructure.Decode(question, &questionResponseDto)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}
		responseDto = append(responseDto, questionResponseDto)
	}
	context.AbortWithStatusJSON(http.StatusOK, responseDto)
}

// option handlers
func CreateOption(context *gin.Context) {

	var optionDto dtos.CreateOptionDto
	if err := context.ShouldBindJSON(&optionDto); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var options []models.Option
	if err := db.Where("question_id = ?", optionDto.QuestionID).Find(&options).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if len(options) >= 3 {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "You can't have more than 3 options per question",
		})
		return
	}
	var checkOptions []models.Option
	for i := range options {
		if options[i].IsCorrect {
			checkOptions = append(checkOptions, options[i])
		}
	}

	if (len(checkOptions) >= 1) && (optionDto.IsCorrect) {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "You can't have more than one correct option per question",
		})
		return
	}

	if (len(checkOptions) < 1) && (len(options) == 2) && (!optionDto.IsCorrect) {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "You can't have all options wrong per question",
		})
		return
	}

	var newOption = models.Option{
		ID:         uuid.New(),
		Text:       optionDto.Text,
		IsCorrect:  optionDto.IsCorrect,
		QuestionID: uuid.MustParse(optionDto.QuestionID),
	}
	if err := db.Create(&newOption).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.AbortWithStatusJSON(http.StatusCreated, newOption)
}

func GetOptions(context *gin.Context) {
	var options []models.Option
	if err := db.Find(&options).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	var responseDto []dtos.OptionResponseDto

	for _, option := range options {
		var optionResponseDto dtos.OptionResponseDto
		err := mapstructure.Decode(option, &optionResponseDto)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}
		responseDto = append(responseDto, optionResponseDto)
	}
	context.AbortWithStatusJSON(http.StatusOK, responseDto)
}

func GetOptionsByQuestionId(context *gin.Context) {
	var options []models.Option
	if err := db.Where("question_id = ?", context.Params.ByName("id")).Find(&options).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Shuffle the questions
	for i := range options {
		j := rand.Intn(i + 1)
		options[i], options[j] = options[j], options[i]
	}

	var responseDto []dtos.OptionResponseDto

	for _, option := range options {
		var optionResponseDto dtos.OptionResponseDto
		err := mapstructure.Decode(option, &optionResponseDto)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}
		responseDto = append(responseDto, optionResponseDto)
	}
	context.AbortWithStatusJSON(http.StatusOK, responseDto)
}
