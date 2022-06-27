package repository

import (
	"context"
)

type Repository interface {
	InsertOneUser(ctx context.Context, document interface{}) (*string, error)
}
