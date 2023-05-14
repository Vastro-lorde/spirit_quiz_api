package handlers

import (
	"net/http"
	dtos "spirit_quiz/internal/Dtos"
	"spirit_quiz/internal/data/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateResult(context *gin.Context) {

	var createResultDto dtos.CreateResultDto
	if err := context.ShouldBindJSON(&createResultDto); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var newResult = models.Result{
		ID:         uuid.New(),
		UserID:     uuid.MustParse(createResultDto.UserID),
		CategoryID: uuid.MustParse(createResultDto.CategoryID),
		Score:      createResultDto.Score,
		Duration:   createResultDto.Duration,
	}

	var user models.User
	if err := db.First(&user, newResult.UserID).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	var category models.Category
	if err := db.First(&category, newResult.CategoryID).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	newResult.Email = user.Email
	newResult.Name = user.Name
	newResult.CategoryName = category.Name

	if err := db.Create(&newResult).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusCreated, newResult)
}

func GetResultsByUserid(context *gin.Context) {
	var results []models.Result
	if err := db.Where("user_id = ?", context.Params.ByName("id")).Order("created_at DESC").Find(&results).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.AbortWithStatusJSON(http.StatusOK, results)
}

func GetResultsByCategoryId(context *gin.Context) {
	var results []models.Result
	if err := db.Where("category_id = ?", context.Params.ByName("id")).Order("score DESC").Find(&results).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.AbortWithStatusJSON(http.StatusOK, results)
}
