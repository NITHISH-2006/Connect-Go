package model

import "time"

type Post struct {
	ID string `json:"id"`
	UserID string `json:"user_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Tags []string `json:"tags"`
	CreatedAt time.Time `json:"create_at"`
	CommentCount int `json:"comment_count"`
}

type CreatedPostRequest struct{
	Title   string   `json:"title" validate:"required"`
    Content string   `json:"content" validate:"required"`
    Tags    []string `json:"tags"`
}

type Comment struct {
	Title string `json:""`
	PostID string `json:""`
	UserID string `json:""`
	Content string `json:""`
	CreatedAt time.Time `json:""`
}

type CreateCommentRequest struct{
	Content string `json:"content" validate:"required"`
}
