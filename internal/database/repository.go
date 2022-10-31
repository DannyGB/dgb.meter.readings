package database

import (
	"context"
	"dgb/meter.readings/internal/configuration"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Repository struct {
	config configuration.Configuration
}

type PageParams struct {
	Skip          int
	Take          int
	SortDirection string
}

func (repository *Repository) GetAll(pageParams PageParams) []primitive.M {

	connect(repository.config)
	coll := repository.getCollection()
	sortDir := 1

	if pageParams.SortDirection == "desc" {
		sortDir = -1
	}

	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"readingdate", sortDir}, {"reading", sortDir}}).SetLimit(int64(pageParams.Take)).SetSkip(int64(pageParams.Skip))
	cursor, err := coll.Find(context.TODO(), filter, opts)

	if err == mongo.ErrNoDocuments {
		return nil
	}

	if err != nil {
		panic(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}

func (repository *Repository) Count() int64 {
	connect(repository.config)
	coll := repository.getCollection()

	filter := bson.D{}
	count, err := coll.CountDocuments(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	return count
}

func (repository *Repository) GetSingle(id string) bson.M {

	connect(repository.config)
	coll := repository.getCollection()

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
	coll := repository.getCollection()

	result, err := coll.InsertOne(context.TODO(), data)

	if err != nil {
		return nil, errors.New("Could not insert document")
	}

	return result.InsertedID, nil
}

func (repository *Repository) Update(id interface{}, data bson.M) error {

	connect(repository.config)
	coll := repository.getCollection()

	filter := bson.D{{"_id", id}}
	_, err := coll.ReplaceOne(context.TODO(), filter, data)

	if err != nil {
		return errors.New("Could not insert document")
	}

	return nil
}

func (repository *Repository) Delete(id interface{}) (deletedCount int, err error) {

	connect(repository.config)
	coll := repository.getCollection()

	filter := bson.D{{"_id", id}}
	result, err := coll.DeleteOne(context.TODO(), filter)

	if result.DeletedCount <= 0 || err != nil {
		return int(result.DeletedCount), errors.New("Could not delete")
	}

	return int(result.DeletedCount), nil
}

func connect(config configuration.Configuration) {

	if client != nil {
		return
	}

	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MONGO_CONNECTION))

	if err != nil {
		panic(err)
	}
}

func (repository *Repository) getCollection() *mongo.Collection {
	return client.Database(repository.config.MONGO_DB).Collection(repository.config.MONGO_COLLECTION)
}

func NewRepository(cfg configuration.Configuration) *Repository {
	return &Repository{
		cfg,
	}
}
