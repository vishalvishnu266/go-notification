package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecieverOption struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	User    string             `bson:"user" json:"user"`
	Webpush bool               `bson:"webpush" json:"webpush"`
	Email   bool               `bson:"email" json:"email"`
	Sms     bool               `bson:"sms" json:"sms"`
}
