package dtos

type UpdateQuestionDto struct {
	CategoryID string `json:"category_id"`
	Text       string `json:"text"`
}
