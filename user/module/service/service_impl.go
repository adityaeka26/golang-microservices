package service

import (
	"context"
	"net/http"

	"github.com/adityaeka26/golang-microservices/user/helper"
	"github.com/adityaeka26/golang-microservices/user/jwt"
	"github.com/adityaeka26/golang-microservices/user/module/model/domain"
	"github.com/adityaeka26/golang-microservices/user/module/model/web"
	"github.com/adityaeka26/golang-microservices/user/module/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type ServiceImpl struct {
	repository repository.Repository
	jwtAuth    jwt.JWT
}

func NewService(repository repository.Repository, jwtAuth jwt.JWT) Service {
	return &ServiceImpl{
		repository: repository,
		jwtAuth:    jwtAuth,
	}
}

func (service *ServiceImpl) CreateUser(ctx context.Context, request web.CreateUserRequest) (*web.CreateUserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	insertedId, err := service.repository.InsertOneUser(ctx, domain.InsertUser{
		Username: request.Username,
		Password: string(hashedPassword),
		Name:     request.Name,
	})
	if err != nil {
		return nil, helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	token, err := service.jwtAuth.GenerateToken(jwt.Payload{
		Id: *insertedId,
	})
	if err != nil {
		helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	return &web.CreateUserResponse{
		Token: *token,
	}, nil
}

func (service *ServiceImpl) GetUser(ctx context.Context, request web.GetUserRequest) (*web.GetUserResponse, error) {
	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	user, err := service.repository.FindOneUser(ctx, bson.M{
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
