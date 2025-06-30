package repository

import (
	"p2-livecode-3-JerSbs/config"
	"p2-livecode-3-JerSbs/models"
)

func GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	err := config.GetDB().Preload("User").Find(&posts).Error
	return posts, err
}
