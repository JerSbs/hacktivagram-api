package repository

import (
	"p2-livecode-3-JerSbs/config"
	"p2-livecode-3-JerSbs/models"
)

func DeletePost(post *models.Post) error {
	return config.GetDB().Delete(post).Error
}
