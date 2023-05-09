package dtos

type RegisterUser struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	ImageUrl string `json:"image_url"`
	Password string `json:"password" binding:"required"`
}
