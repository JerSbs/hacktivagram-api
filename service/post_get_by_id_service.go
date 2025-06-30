package service

import (
	"p2-livecode-3-JerSbs/dto"
	"p2-livecode-3-JerSbs/repository"

	"gorm.io/gorm"
)

func GetPostByIDService(id uint) (*dto.PostDetailResponse, error) {
	post, err := repository.GetPostByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, ErrInternal
	}

	var comments []dto.CommentItem
	for _, c := range post.Comments {
		comments = append(comments, dto.CommentItem{
			ID:      c.ID,
			Content: c.Content,
			UserID:  c.UserID,
			Name:    c.User.FullName,
		})
	}

	response := &dto.PostDetailResponse{
		ID:       post.ID,
		Content:  post.Content,
		ImageURL: post.ImageURL,
		UserID:   post.UserID,
		UserName: post.User.FullName,
		Comments: comments,
	}

	return response, nil
}
