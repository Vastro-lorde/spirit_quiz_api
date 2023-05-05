package handlers

import (
	"net/http"
	dtos "spirit_quiz/internal/Dtos"
	"spirit_quiz/internal/data/models"
	"spirit_quiz/internal/middlewares"
	"spirit_quiz/internal/services"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Sign up handler
func SignUp(context *gin.Context) {

	var registerDto dtos.RegisterUser
	if err := context.ShouldBindJSON(&registerDto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if strings.Index(registerDto.Email, "+") != -1 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Email address cannot contain a plus sign"})
		context.Abort()
		return
	}

	hashedPassword, err := services.HashPassword(registerDto.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	var newUser = models.User{
		ID:       uuid.New(),
		Name:     registerDto.FullName,
		Email:    registerDto.Email,
		Password: hashedPassword,
	}

	if strings.ToLower(registerDto.Email) == "omatsolaseund@gmail.com" {
		newUser.Role = "ADMIN"
	}

	if err := db.Create(&newUser).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	alphaNumericToken := services.GenerateRandomAlphaNumericString(6)

	if err := services.SendWelcomeEmail(newUser.Email, alphaNumericToken); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, newUser)
	context.Abort()
}

func Login(context *gin.Context) {
	var loginDto dtos.LoginDto
	if err := context.ShouldBindJSON(&loginDto); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := db.Where("email = ?", loginDto.Email).First(&user).Error.Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	check := services.ComparePassword(loginDto.Password, user.Password)
	if !check {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "wrong password"})
		return
	}

	var userClaim dtos.UserClaim
	userClaim.Email = user.Email
	userClaim.ID = user.ID.String()
	userClaim.Role = string(user.Role)

	token, err := middlewares.GenerateToken(userClaim)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.AbortWithStatusJSON(http.StatusOK, gin.H{
		"token": token,
		"exp":   time.Now().Add(time.Hour * 24 * 6).Unix(),
		"user": userClaim,
	})
}


func RequestResetPassword(context *gin.Context)  {
	var email string
	if err := context.ShouldBindJSON(&email); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error.Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	otp := services.GenerateRandomNumbers(5)

	hashedOtp, err := services.HashPassword(otp)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	
	// send the otp to the user email
	if err := services.SendPasswordResetEmail(user.Email, otp); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update the user with the generated OTP
	user.Otp = hashedOtp

	if err := db.Save(&user).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Return the updated user object
	context.AbortWithStatusJSON(http.StatusOK, gin.H{"success": "an email has been sent to "+ email})
}

func ResetPassword(context *gin.Context)  {
	var resetPasswordDto dtos.ResetPasswordDto
	if err := context.ShouldBindJSON(&resetPasswordDto); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Where("email = ?", resetPasswordDto.Email).First(&user).Error.Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	check := services.ComparePassword(resetPasswordDto.Otp, user.Otp)
	if !check {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid otp"})
		return
	}

	// send the otp to the user email
	if err := services.SendResetPasswordSuccessEmail(user.Email); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update the user with the generated OTP
	user.Otp = ""
	user.Password = resetPasswordDto.Password

	if err := db.Save(&user).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	user.Password = ""
	// Return the updated user object
	context.AbortWithStatusJSON(http.StatusOK, gin.H{"user": user})

}