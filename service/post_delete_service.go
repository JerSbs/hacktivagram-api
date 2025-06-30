package service

import (
	"fmt"
	"p2-livecode-3-JerSbs/models"
	"p2-livecode-3-JerSbs/repository"

	"gorm.io/gorm"
)

func DeletePostService(postID, userID uint) (*models.Post, error) {
	post, err := repository.GetPostByID(postID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, ErrInternal
	}

	if post.UserID != userID {
		return nil, ErrForbidden
	}

	if err := repository.DeletePost(post); err != nil {
		return nil, ErrInternal
	}

	// Log
	desc := fmt.Sprintf("user deleted POST with ID %d", post.ID)
	_ = repository.CreateUserLog(userID, desc)

	return post, nil
}
