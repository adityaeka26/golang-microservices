package repository

import (
	"context"

	"github.com/adityaeka26/golang-microservices/user/module/model/domain"
)

type Repository interface {
	InsertOneUser(ctx context.Context, document interface{}) (*string, error)
	FindOneUser(ctx context.Context, filter interface{}) (*domain.User, error)
}
