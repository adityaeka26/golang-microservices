package service

import (
	"context"
	"net/http"

	"github.com/adityaeka26/golang-microservices/user/helper"
	"github.com/adityaeka26/golang-microservices/user/module/model/domain"
	"github.com/adityaeka26/golang-microservices/user/module/model/web"
	"github.com/adityaeka26/golang-microservices/user/module/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (service *ServiceImpl) GetUser(ctx context.Context, request web.GetUserRequest) (*web.GetUserResponse, error) {
	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	user, err := service.Repository.FindOneUser(ctx, bson.M{
		"_id": id,
	})
	if err != nil {
		return nil, helper.CustomError(http.StatusInternalServerError, err.Error())
	}
	if user == nil {
		return nil, helper.CustomError(http.StatusNotFound, "User not found")
	}

	return &web.GetUserResponse{
		Id:       user.Id,
		Username: user.Username,
		Name:     user.Name,
	}, nil
}
