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
	"github.com/mitchellh/mapstructure"
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
		Email:    strings.ToLower(registerDto.Email),
		ImageUrl: registerDto.ImageUrl,
		Password: hashedPassword,
	}

	if context.Params.ByName("google") == "true" {
		newUser.Verified = true
	}

	if strings.ToLower(registerDto.Email) == "omatsolaseund@gmail.com" {
		newUser.Role = "ADMIN"
	}

	alphaNumericToken := services.GenerateRandomAlphaNumericString(6)
	hashedOtp, err := services.HashPassword(alphaNumericToken)

	if err := services.SendWelcomeEmail(newUser.Email, alphaNumericToken); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	newUser.Otp = hashedOtp

	if err := db.Create(&newUser).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	newUser.Password = ""
	newUser.Otp = ""

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

	if err := db.Where("email = ?", strings.ToLower(loginDto.Email)).First(&user).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !user.Verified {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user email unverified"})
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

	var userResponseDto dtos.UserResponseDto
	err = mapstructure.Decode(user, &userResponseDto)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.AbortWithStatusJSON(http.StatusOK, gin.H{
		"token":     token,
		"exp":       time.Now().Add(time.Hour * 24 * 6).Unix(),
		"userClaim": userClaim,
		"user":      userResponseDto,
	})
}

func RequestResetPassword(context *gin.Context) {

	type EmailDto struct {
		Email string `json:"email" binding:"required"`
	}
	var email EmailDto
	if err := context.ShouldBindJSON(&email); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Where("email = ?", email.Email).First(&user).Error; err != nil {
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
	context.AbortWithStatusJSON(http.StatusOK, gin.H{"success": "an email has been sent to " + email.Email})
}

func ResetPassword(context *gin.Context) {
	var resetPasswordDto dtos.ResetPasswordDto
	if err := context.ShouldBindJSON(&resetPasswordDto); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Where("email = ?", resetPasswordDto.Email).First(&user).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	hashedPassword, err := services.HashPassword(resetPasswordDto.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	user.Otp = ""
	user.Password = hashedPassword

	if err := db.Save(&user).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = ""
	// Return the updated user object
	context.AbortWithStatusJSON(http.StatusOK, gin.H{
		"success": "password has been reset",
		"user":    user,
	})
}

func VerifyEmail(context *gin.Context) {
	var verifyDto dtos.VerifyEmailDto
	if err := context.ShouldBindJSON(&verifyDto); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Where("email = ?", verifyDto.Email).First(&user).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	check := services.ComparePassword(verifyDto.Otp, user.Otp)
	if !check {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid otp"})
		return
	}

	// send the otp to the user email
	if err := services.SendSuccessEmailVerification(user.Email); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update the user with the generated OTP
	user.Otp = ""
	user.Verified = true

	if err := db.Save(&user).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	user.Password = ""
	// Return the updated user object
	context.AbortWithStatusJSON(http.StatusOK, gin.H{"user": user})
}
