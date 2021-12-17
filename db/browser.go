package db

import (
	"context"

	"github.com/growmax/noti/model"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
)

func AddBrowser(browser model.Browser) error {
	//Get MongoDB connection using connectionhelper.
	client, err := GetMongoClient()
	if err != nil {
		return err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection("browser")
	//Perform InsertOne operation & validate against the error.
	_, err = collection.InsertOne(context.TODO(), browser)
	if err != nil {
		return err
	}
	//Return success without any error.
	return nil
}
