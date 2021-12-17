package model

import (
	webpush "github.com/SherClockHolmes/webpush-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Browser struct {
	ID      primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	User    string               `bson:"user" json:"user,omitempty"`
	Browser webpush.Subscription `bson:"browser" json:"browser,omitempty"`
}
