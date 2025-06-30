package repository

import (
	"p2-livecode-3-JerSbs/config"
	"p2-livecode-3-JerSbs/models"
)

func GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	err := config.GetDB().
		Preload("User").
		Preload("Comments.User").
		First(&post, id).Error

	if err != nil {
		return nil, err
	}
	return &post, nil
}
