package service

import (
	"context"

	"github.com/adityaeka26/golang-microservices/user/module/model/web"
)

type Service interface {
	CreateUser(ctx context.Context, request web.CreateUserRequest) (*web.CreateUserResponse, error)
	GetUser(ctx context.Context, request web.GetUserRequest) (*web.GetUserResponse, error)
}
