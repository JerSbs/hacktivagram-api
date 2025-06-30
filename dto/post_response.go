package dto

type PostResponse struct {
	ID       uint   `json:"id"`
	Content  string `json:"content"`
	ImageURL string `json:"image_url"`
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
}
