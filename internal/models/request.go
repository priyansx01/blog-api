package models

type CreatePostRequest struct {
	Title    string   `json:"title" binding:"requird,min=1,max-200"`
	Content  string   `json:"content" binding:"required, min=1"`
	Category string   `json:"category" binding:"required"`
	Tags     []string `json:"tags"`
}

type UpdatePostRequest struct {
	Title    *string  `json:"title" binding:"omitempty,min=1,max=200"`
	Content  *string  `json:"content" binding:"omitempty,min=1"`
	Category *string  `json:"category" binding:"omitempty"`
	Tags     []string `json:"tags"`
}
