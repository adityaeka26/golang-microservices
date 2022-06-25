package service

import (
	"context"

	"github.com/adityaeka26/golang-microservices/user/module/model/web"
)

type Service interface {
	CreateUser(ctx context.Context, request web.CreateUserRequest) error
}
