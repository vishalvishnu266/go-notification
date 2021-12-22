package db

import (
	"context"
	// "time"

	"github.com/growmax/noti/model"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func CreateUser(user model.User) error {
	client, err := GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(DB).Collection("user")
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func IsUserExist(user string) (bool, error) {
	client, err := GetMongoClient()
	if err != nil {
		return false, err
	}
	collection := client.Database(DB).Collection("user")
	filter := bson.M{"user": user}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
