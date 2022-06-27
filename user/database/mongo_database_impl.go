package database

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adityaeka26/golang-microservices/user/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDatabaseImpl struct {
	MongoClient  *mongo.Client
	DatabaseName string
}

func NewMongoDB(uri string, dbName string) MongoDatabase {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	return &MongoDatabaseImpl{
		MongoClient:  client,
		DatabaseName: dbName,
	}
}

type InsertOne struct {
	CollectionName string
	Document       interface{}
}

func (mongoDatabase MongoDatabaseImpl) InsertOne(ctx context.Context, payload InsertOne) (*string, error) {
	collection := mongoDatabase.MongoClient.Database(mongoDatabase.DatabaseName).Collection(payload.CollectionName)
	insertDoc, err := collection.InsertOne(ctx, payload.Document)
	if err != nil {
		return nil, helper.CustomError(http.StatusInternalServerError, fmt.Sprintf("Error Mongodb Connection: %s", err.Error()))
	}

	insertedId := insertDoc.InsertedID.(primitive.ObjectID).Hex()
	return &insertedId, nil
}
