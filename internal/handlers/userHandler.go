package handlers

import (
	"fmt"
	"net/http"
	dtos "spirit_quiz/internal/Dtos"
	"spirit_quiz/internal/data/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

func UpdateUser(context *gin.Context) {

	var updateUser dtos.UpdateUserDto
	if err := context.ShouldBindJSON(&updateUser); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	if err := db.Where("id = ?", updateUser.Id).Find(&user).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println(user)

	if updateUser.Name != "" {
		user.Name = updateUser.Name
	}
	if updateUser.ImageUrl != "" {
		user.ImageUrl = updateUser.ImageUrl
	}

	if err := db.Save(&user).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var userResponseDto dtos.UserResponseDto
	err := mapstructure.Decode(user, &userResponseDto)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.AbortWithStatusJSON(http.StatusOK, gin.H{
		"message": "successfully updated",
		"data":    userResponseDto,
	})
}

func GetUserById(context *gin.Context) {
	var user models.User
	// var getUser struct {
	// 	ID uuid.UUID `json:"id"`
	// }
	// if err := context.ShouldBindJSON(&getUser); err != nil {
	// 	context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	// 	return
	// }

	if err := db.Where("id = ?", context.Params.ByName("id")).First(&user).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	var userResponseDto dtos.UserResponseDto
	err := mapstructure.Decode(user, &userResponseDto)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.AbortWithStatusJSON(http.StatusOK, userResponseDto)
}

func GetUsers(context *gin.Context) {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	var responseDto []dtos.UserResponseDto

	for _, user := range users {
		var userResponseDto dtos.UserResponseDto
		err := mapstructure.Decode(&user, &userResponseDto)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}
		userResponseDto.CreatedAt = user.CreatedAt
		userResponseDto.UpdatedAt = user.UpdatedAt
		responseDto = append(responseDto, userResponseDto)
	}

	context.AbortWithStatusJSON(http.StatusOK, responseDto)
}

func DeleteUserById(context *gin.Context) {

	var user models.User
	var getUser struct {
		ID uuid.UUID `json:"id"`
	}
	if err := context.ShouldBindJSON(&getUser); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := db.First(&user, getUser.ID).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.AbortWithStatusJSON(http.StatusOK, gin.H{
		"message": "successfully deleted",
		"data":    user,
	})
}
