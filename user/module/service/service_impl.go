package service

import (
	"context"
	"net/http"

	"github.com/adityaeka26/golang-microservices/user/helper"
	"github.com/adityaeka26/golang-microservices/user/module/model/domain"
	"github.com/adityaeka26/golang-microservices/user/module/model/web"
	"github.com/adityaeka26/golang-microservices/user/module/repository"
	"golang.org/x/crypto/bcrypt"
)

type ServiceImpl struct {
	Repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return &ServiceImpl{
		Repository: repository,
	}
}

func (service *ServiceImpl) CreateUser(ctx context.Context, request web.CreateUserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	_, err = service.Repository.InsertOneUser(ctx, domain.InsertUser{
		Username: request.Username,
		Password: string(hashedPassword),
		Name:     request.Name,
	})
	if err != nil {
		return helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
