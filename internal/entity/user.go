package entity

import (
	"time"

	"gorm.io/gorm"
)

// Userエンティティ
type User struct {
	ID        string         `gorm:"primaryKey"`
	AuthID    string         `gorm:"unique;not null"`
	Name      string         `gorm:"not null"`
	Email     string         `gorm:"unique;not null"`
	Role      string         `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
}
