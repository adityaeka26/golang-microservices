package database

import "context"

type MongoDatabase interface {
	InsertOne(ctx context.Context, payload InsertOne) (*string, error)
	FindOne(ctx context.Context, payload FindOne) error
}
