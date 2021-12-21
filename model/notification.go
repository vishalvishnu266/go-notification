package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	User         string             `bson:"user" json:"-"`
	Notification string             `bson:"notification" json:"notification"`
	CreatedAt    time.Time          `bson:"created_at" json:"time"`
}
