package dtos

type CreateQuestionDto struct {
	CategoryID string `json:"category_id" binding:"required"`
	Text       string `json:"text" binding:"required"`
}
