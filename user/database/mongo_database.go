package database

import "context"

type MongoDatabase interface {
	InsertOne(ctx context.Context, payload InsertOne) (*string, error)
}
