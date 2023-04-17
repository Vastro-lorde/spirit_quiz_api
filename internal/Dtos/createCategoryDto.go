package dtos

type CreateCategoryDto struct{
	Name 		string `json:"name" binding:"required"`
}