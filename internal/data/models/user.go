package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	ADMIN       Role = "ADMIN"
	CLIENT      Role = "CLIENT"
	DefaultRole      = CLIENT
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"not null;uniqueIndex"`
	Password  string    `gorm:"not null"`
	Otp       string    `gorm:"not null;default:''"`
	Role      Role      `gorm:"not null;default:CLIENT"`
	Results   []Result  `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
