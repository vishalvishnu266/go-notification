package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Browser struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	User      string             `bson:"user" json:"user"`
	Browser   string             `bson:"browser" json:"browser"`
	CreatedAt time.Time          `bson:"created_at" json:"time"`
}
