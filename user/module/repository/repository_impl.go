package repository

import (
	"context"

	"github.com/adityaeka26/golang-microservices/user/database"
)

type RepositoryImpl struct {
	MongoDatabase database.MongoDatabase
}

func NewRepository(mongoDatabase database.MongoDatabase) Repository {
	return &RepositoryImpl{
		MongoDatabase: mongoDatabase,
	}
}

func (repository *RepositoryImpl) InsertOneUser(ctx context.Context, document interface{}) (*string, error) {
	insertedId, err := repository.MongoDatabase.InsertOne(ctx, database.InsertOne{
		CollectionName: "user",
		Document:       document,
	})
	if err != nil {
		return nil, err
	}

	return insertedId, nil
}
