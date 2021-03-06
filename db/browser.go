package db

import (
	"context"

	"github.com/growmax/noti/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddBrowser(browser model.Browser) error {
	client, err := GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(DB).Collection("browser")
	_, err = collection.InsertOne(context.TODO(), browser)
	if err != nil {
		return err
	}
	return nil
}

func GetBrowser(user string, page int) ([]model.Browser, error) {
	var browsers []model.Browser
	client, err := GetMongoClient()
	if err != nil {
		return browsers, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection("browser")
	filter := bson.M{"user": user}
	findOptions := options.Find()
	var perPage int64 = 3
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ctx := context.TODO()

	// total, _ := collection.CountDocuments(ctx, filter)
	findOptions.SetSort(bson.D{{"created_at", -1}})

	findOptions.SetSkip((int64(page) - 1) * perPage)
	findOptions.SetLimit(perPage)

	cursor, _ := collection.Find(ctx, filter, findOptions)
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var browser model.Browser
		cursor.Decode(&browser)
		browsers = append(browsers, browser)
	}
	return browsers, nil
}
