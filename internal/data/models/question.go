package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	Text       string    `gorm:"not null"`
	Options    []Option  `gorm:"foreignKey:QuestionID"`
	CategoryID uuid.UUID `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
