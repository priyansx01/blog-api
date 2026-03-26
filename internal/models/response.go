package models

import (
	"time"
)

type PostResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ErrorResponse struct {
	Error   string                 `json:"error"`
	Message string                 `json:"message,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
}

type PostListResponse struct {
	Posts  []PostResponse `json:"posts"`
	Total  int64          `json:"total"`
	Limt   int            `json:"limt"`
	Offset int            `json:"offset"`
}

func ToPostResponse(post *Post) *PostResponse {
	return &PostResponse{
		ID:        post.ID.Hex(),
		Title:     post.Title,
		Content:   post.Content,
		Category:  post.Category,
		Tags:      post.Tags,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

func ToPostResponseList(posts []Post) []PostResponse {
	responses := make([]PostResponse, len(posts))
	for i, post := range posts {
		responses[i] = *ToPostResponse(&post)
	}
	return responses
}
