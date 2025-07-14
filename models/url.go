package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type URL struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OriginalURL string             `bson:"original_url" json:"original_url"`
	ShortCode   string             `bson:"short_code" json:"short_code"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	ExpiresAt   time.Time          `bson:"expires_at" json:"expires_at"`
	ClickCount  int                `bson:"click_count" json:"click_count"`
}
