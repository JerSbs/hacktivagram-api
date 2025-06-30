package repository

import (
	"p2-livecode-3-JerSbs/config"
	"p2-livecode-3-JerSbs/models"
)

func CreatePost(post *models.Post) error {
	return config.GetDB().Create(post).Error
}
