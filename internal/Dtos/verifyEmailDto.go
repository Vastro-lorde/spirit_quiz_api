package dtos

type VerifyEmailDto struct {
	Email    string `json:"email" binding:"required"`
	Otp      string `json:"otp" binding:"required"`
}
