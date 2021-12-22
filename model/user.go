package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	User      string             `bson:"user" json:"user"`
	Tenantid  string             `bson:"browser" json:"browser"`
	IsSeller  bool             `bson:"isseller" json:"isseller"`
	PriEmail  string             `bson:"priemail" json:"priemail"`
	SecEmail  string             `bson:"secemail" json:"secemail"`
	PriPhone  string             `bson:"priphone" json:"priphone"`
	SecPhone  string             `bson:"secphone" json:"secphone"`
	CreatedAt time.Time          `bson:"created_at" json:"time"`
}
