package dto

type CreatePostPayload struct {
	Content string `json:"content" validate:"required"`
	Title   string `json:"title"  validate:"required"`
	UserID  string `json:"user_id"  validate:"required"`
}

type GetPostByUserIDPayload struct {
	UserID     string `json:"user_id"  validate:"required"`
	LastPostID int
	PageLength int
}
