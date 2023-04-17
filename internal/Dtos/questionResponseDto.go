package dtos

import (
	"time"

	"github.com/google/uuid"
)

type QuestionResponseDto struct {
	ID         uuid.UUID `json:"id"`
	Text       string    `json:"text"`
	CategoryID uuid.UUID `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
