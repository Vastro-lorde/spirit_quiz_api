package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Option struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	Text       string    `gorm:"not null"`
	IsCorrect  bool      `gorm:"not null"`
	QuestionID uuid.UUID `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
