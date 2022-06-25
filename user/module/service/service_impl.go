package service

import (
	"context"

	"github.com/adityaeka26/golang-microservices/user/module/model/web"
)

type ServiceImpl struct{}

func NewService() Service {
	return &ServiceImpl{}
}

func (service *ServiceImpl) CreateUser(ctx context.Context, request web.CreateUserRequest) error {
	return nil
}
