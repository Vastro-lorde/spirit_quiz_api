package dtos

import (
	"spirit_quiz/internal/data/models"
	"time"
)

type UserResponseDto struct {
	ID        string      `json:"id"`
	FullName  string      `json:"full_name"`
	Email     string      `json:"email"`
	ImageUrl  string      `json:"image_url"`
	Role      models.Role `json:"role"`
	Verified  bool        `json:"verified"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
