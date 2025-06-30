package models

import "gorm.io/gorm"

type UserActivityLog struct {
	gorm.Model
	Description string `gorm:"not null" json:"description"`
	UserID      uint   `gorm:"not null" json:"user_id"`
	User        User   `gorm:"foreignKey:UserID"`
}
