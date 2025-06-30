package dto

type CreatePostRequest struct {
	Content  string `json:"content"`
	ImageURL string `json:"image_url" validate:"required,url"`
}
