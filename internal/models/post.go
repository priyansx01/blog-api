package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Post struct {
	ID        bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string        `json:"title" bson:"title"`
	Content   string        `json:"content"  bson:"content"`
	Category  string        `json:"category" bson:"category"`
	Tags      []string      `json:"tags" bson:"tags"`
	CreatedAt time.Time     `json:"createdAt"  bson:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"  bson:"updatedAt"`
}
