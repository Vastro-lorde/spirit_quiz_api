package dtos

type ResetPasswordDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Otp      string `json:"otp" binding:"required"`
}
