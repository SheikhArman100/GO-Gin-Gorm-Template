package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Email     string         `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-"`
}
