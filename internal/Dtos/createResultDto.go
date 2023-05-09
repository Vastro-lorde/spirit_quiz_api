package dtos

type CreateResultDto struct {
	CategoryID string `json:"category_id" binding:"required"`
	UserID     string `json:"user_id" binding:"required"`
	Score      string `json:"score" binding:"required"`
}
