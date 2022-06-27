package repository

import (
	"context"

	"github.com/adityaeka26/golang-microservices/user/database"
	"github.com/adityaeka26/golang-microservices/user/module/model/domain"
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

func (repository *RepositoryImpl) FindOneUser(ctx context.Context, filter interface{}) (*domain.User, error) {
	var result *domain.User
	err := repository.MongoDatabase.FindOne(ctx, database.FindOne{
		CollectionName: "user",
		Filter:         filter,
		Result:         &result,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
