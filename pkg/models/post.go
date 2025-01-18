package models

import "time"

type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    string   `json:"user_id"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	Tags      []string `json:"tags"`
}

func NewPost(ID int64, content, title, userID string, tags []string) *Post {
	return &Post{
		Content:   content,
		Title:     title,
		UserID:    userID,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
		Tags:      tags,
		ID:        ID,
	}
}
