package service

import (
	"p2-livecode-3-JerSbs/dto"
	"p2-livecode-3-JerSbs/repository"
)

func GetAllPostsService() ([]dto.PostResponse, error) {
	posts, err := repository.GetAllPosts()
	if err != nil {
		return nil, ErrInternal
	}

	var result []dto.PostResponse
	for _, p := range posts {
		result = append(result, dto.PostResponse{
			ID:       p.ID,
			Content:  p.Content,
			ImageURL: p.ImageURL,
			UserID:   p.UserID,
			UserName: p.User.FullName,
		})
	}

	return result, nil
}
