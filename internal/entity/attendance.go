package entity

import (
	"time"
)

// 勤怠エンティティ
type Attendance struct {
	ID        string    `gorm:"primaryKey"`
	UserID    string    `gorm:"primaryKey"`
	Date      time.Time `gorm:"primaryKey"`
	CheckIn   time.Time
	CheckOut  time.Time
	IsLate    bool
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
