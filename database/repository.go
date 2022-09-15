package database

import (
	"context"
	"dgb/meter.readings/application"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func GetReading(id string, cfg application.Configuration) bson.M {

	connect(cfg)

	coll := client.Database("sample_airbnb").Collection("listingsAndReviews")

	filter := bson.M{"_id": id}

	var result bson.M
	res := coll.FindOne(context.TODO(), filter)
	err := res.Decode(&result)

	if err == mongo.ErrNoDocuments {
		return nil
	}

	if err != nil {
		panic(err)
	}

	return result
}

func connect(cfg application.Configuration) {

	if client != nil {
		return
	}

	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.MONGO_CONNECTION))

	if err != nil {
		panic(err)
	}
}
