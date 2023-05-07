package dtos

type UpdateOptionDto struct {
	QuestionID string `json:"question_id"`
	Text       string `json:"text"`
	IsCorrect  bool   `json:"is_correct"`
}
