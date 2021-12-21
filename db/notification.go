package db

import (
	"context"
	// "time"

	"github.com/growmax/noti/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddNotification(notification model.Notification) error {
	client, err := GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(DB).Collection("notification")
	_, err = collection.InsertOne(context.TODO(), notification)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}

func GetNotification(user string, page int) ([]model.Notification, error) {
	var notifications []model.Notification
	client, err := GetMongoClient()
	if err != nil {
		return notifications, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection("notification")
	filter := bson.M{"user": user}
	findOptions := options.Find()
	var perPage int64 = 10
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ctx := context.TODO()

	// total, _ := collection.CountDocuments(ctx, filter)
	findOptions.SetSort(bson.D{{"created_at", -1}})

	findOptions.SetSkip((int64(page) - 1) * perPage)
	findOptions.SetLimit(perPage)

	cursor, _ := collection.Find(ctx, filter, findOptions)
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var notification model.Notification
		cursor.Decode(&notification)
		notifications = append(notifications, notification)
	}
	return notifications, nil
}
