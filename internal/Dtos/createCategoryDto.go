package dtos

type CreateCategoryDto struct {
	Name     string `json:"name" binding:"required"`
	ImageUrl string `json:"image_url"`
}

type UpdateCategoryDto struct {
	Name     string `json:"name" binding:"required"`
	ImageUrl string `json:"image_url"`
}
