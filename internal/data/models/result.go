package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Result struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Score     string    `gorm:"not null"`
	Email     string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	UserID    uuid.UUID `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
