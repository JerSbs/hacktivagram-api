package repository

import (
	"p2-livecode-3-JerSbs/config"
	"p2-livecode-3-JerSbs/models"
)

func CreateUserLog(userID uint, description string) error {
	log := models.UserActivityLog{
		UserID:      userID,
		Description: description,
	}
	return config.GetDB().Create(&log).Error
}
