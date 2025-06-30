package dto

type CommentItem struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
	Name    string `json:"name"`
}

type PostDetailResponse struct {
	ID       uint          `json:"id"`
	Content  string        `json:"content"`
	ImageURL string        `json:"image_url"`
	UserID   uint          `json:"user_id"`
	UserName string        `json:"user_name"`
	Comments []CommentItem `json:"comments"`
}
