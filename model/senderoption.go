package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SenderOption struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	User    string             `bson:"user"`
	Webpush bool               `bson:"webpush"`
	Email   bool               `bson:"email"`
	Sms     bool               `bson:"sms"`
}
