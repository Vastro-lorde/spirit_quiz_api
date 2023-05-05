package dtos

type CreateOptionDto struct {
	QuestionID string `json:"question_id" binding:"required"`
	Text       string `json:"text" binding:"required"`
	IsCorrect  bool   `json:"is_correct" binding:"required"`
}
