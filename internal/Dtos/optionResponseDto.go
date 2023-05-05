package dtos

import (
	"time"

	"github.com/google/uuid"
)

type OptionResponseDto struct {
	ID         uuid.UUID `json:"id"`
	Text       string    `json:"text"`
	IsCorrect  bool      `json:"is_correct"`
	QuestionID uuid.UUID `json:"question_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
