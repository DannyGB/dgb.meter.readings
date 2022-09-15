package database

import (
	"context"
	"dgb/meter.readings/application"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Repository struct {
	config application.Configuration
}

func (repository *Repository) Get(id string) bson.M {

	connect(repository.config)

	coll := client.Database(repository.config.MONGO_DB).Collection(repository.config.MONGO_COLLECTION)

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

func (repository *Repository) Insert(data bson.M) (id interface{}, err error) {

	connect(repository.config)

	coll := client.Database(repository.config.MONGO_DB).Collection(repository.config.MONGO_COLLECTION)
	result, err := coll.InsertOne(context.TODO(), data)

	if err != nil {
		return nil, errors.New("Could not insert document")
	}

	return result.InsertedID, nil
}

func (repository *Repository) Update(id interface{}, data bson.M) error {

	connect(repository.config)

	coll := client.Database(repository.config.MONGO_DB).Collection(repository.config.MONGO_COLLECTION)
	filter := bson.D{{"_id", id}}
	_, err := coll.ReplaceOne(context.TODO(), filter, data)

	if err != nil {
		return errors.New("Could not insert document")
	}

	return nil
}

func (repository *Repository) Delete(id interface{}) (deletedCount int, err error) {

	connect(repository.config)

	coll := client.Database(repository.config.MONGO_DB).Collection(repository.config.MONGO_COLLECTION)
	filter := bson.D{{"_id", id}}
	result, err := coll.DeleteOne(context.TODO(), filter)

	if result.DeletedCount < 0 || err != nil {
		return int(result.DeletedCount), errors.New("Could not delete")
	}

	return int(result.DeletedCount), nil
}

func connect(config application.Configuration) {

	if client != nil {
		return
	}

	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MONGO_CONNECTION))

	if err != nil {
		panic(err)
	}
}

func NewRepository(cfg application.Configuration) *Repository {
	return &Repository{
		cfg,
	}
}
