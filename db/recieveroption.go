package db

import (
	"context"
	"github.com/growmax/noti/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddRecieverOption(r model.RecieverOption) error {
	client, err := GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(DB).Collection("recieveroption")
	_, err = collection.InsertOne(context.TODO(), r)
	if err != nil {
		return err
	}
	return nil
}
func GetRecieverOption(user string) (model.RecieverOption, error) {
	var r model.RecieverOption
	client, err := GetMongoClient()
	if err != nil {
		return r, err
	}
	collection := client.Database(DB).Collection("recieveroption")
	filter := bson.M{"user": user}
	ctx := context.TODO()

	err = collection.FindOne(ctx, filter).Decode(&r)
	if err != nil {

		return r, err
	}
	return r, nil

}

func UpdateWebPush(user string) error {
	client, err := GetMongoClient()
	if err != nil {
		return err
	}
	opts := options.Update().SetUpsert(true)

	collection := client.Database(DB).Collection("recieveroption")
	filter := bson.M{"user": user}
	val1 := bson.M{"$not": "$webpush"}
	val2 := bson.M{"webpush": val1}
	val := bson.M{"$set": val2}
	option := bson.A{val}

	_, _ = collection.UpdateOne(
		context.Background(),
		filter,
		option, opts)
	return nil
}

func UpdateSms(user string) error {
	client, err := GetMongoClient()
	if err != nil {
		return err
	}
	opts := options.Update().SetUpsert(true)

	collection := client.Database(DB).Collection("recieveroption")
	filter := bson.M{"user": user}
	val1 := bson.M{"$not": "$sms"}
	val2 := bson.M{"sms": val1}
	val := bson.M{"$set": val2}
	option := bson.A{val}

	_, _ = collection.UpdateOne(
		context.Background(),
		filter,
		option, opts)
	return nil
}

func UpdateEmail(user string) error {
	client, err := GetMongoClient()
	if err != nil {
		return err
	}
	opts := options.Update().SetUpsert(true)

	collection := client.Database(DB).Collection("recieveroption")
	filter := bson.M{"user": user}
	val1 := bson.M{"$not": "$email"}
	val2 := bson.M{"email": val1}
	val := bson.M{"$set": val2}
	option := bson.A{val}

	_, _ = collection.UpdateOne(
		context.Background(),
		filter,
		option, opts)
	return nil
}
