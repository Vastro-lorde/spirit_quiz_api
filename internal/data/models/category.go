package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Name      string     `gorm:"not null"`
	ImageUrl  string     `gorm:"not null;default:''"`
	Questions []Question `gorm:"foreignKey:CategoryID"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}
