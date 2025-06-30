package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Content  string    `gorm:"not null" json:"content"`
	ImageURL string    `gorm:"not null" json:"image_url"`
	UserID   uint      `gorm:"not null" json:"user_id"`
	User     User      `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"constraint:OnDelete:CASCADE"`
}
