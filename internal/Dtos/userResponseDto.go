package dtos

import (
	"spirit_quiz/internal/data/models"
	"time"

	"github.com/google/uuid"
)

type UserResponseDto struct {
	ID        uuid.UUID   `json:"id"`
	FullName  string      `json:"full_name"`
	Email     string      `json:"email"`
	ImageUrl  string      `json:"image_url"`
	Role      models.Role `json:"role"`
	Verified  bool        `json:"verified"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
