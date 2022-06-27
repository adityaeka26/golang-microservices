package repository

import (
	"context"

	"github.com/adityaeka26/golang-microservices/user/database"
	"github.com/adityaeka26/golang-microservices/user/module/model/domain"
)

type RepositoryImpl struct {
	mongoDatabase database.MongoDatabase
}

func NewRepository(mongoDatabase database.MongoDatabase) Repository {
	return &RepositoryImpl{
		mongoDatabase: mongoDatabase,
	}
}

func (repository *RepositoryImpl) InsertOneUser(ctx context.Context, document interface{}) (*string, error) {
	insertedId, err := repository.mongoDatabase.InsertOne(ctx, database.InsertOne{
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
	err := repository.mongoDatabase.FindOne(ctx, database.FindOne{
		CollectionName: "user",
		Filter:         filter,
		Result:         &result,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
