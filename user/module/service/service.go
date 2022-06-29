package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/adityaeka26/golang-microservices/user/database"
	"github.com/adityaeka26/golang-microservices/user/helper"
	"github.com/adityaeka26/golang-microservices/user/jwt"
	"github.com/adityaeka26/golang-microservices/user/kafka"
	"github.com/adityaeka26/golang-microservices/user/module/model/domain"
	"github.com/adityaeka26/golang-microservices/user/module/model/web"
	"github.com/adityaeka26/golang-microservices/user/module/repository"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(ctx context.Context, request web.RegisterRequest) error
	VerifyRegister(ctx context.Context, request web.VerifyRegisterRequest) (*web.VerifyRegisterResponse, error)
}
type ServiceImpl struct {
	repository    repository.Repository
	jwtAuth       jwt.JWT
	redis         database.Redis
	kafkaProducer kafka.KafkaProducer
}

func NewService(repository repository.Repository, jwtAuth jwt.JWT, redis database.Redis, kafkaProducer kafka.KafkaProducer) Service {
	return &ServiceImpl{
		repository:    repository,
		jwtAuth:       jwtAuth,
		redis:         redis,
		kafkaProducer: kafkaProducer,
	}
}

func (service ServiceImpl) Register(ctx context.Context, request web.RegisterRequest) error {
	// Preprocecss username & mobile number
	request.Username = strings.ToLower(request.Username)
	mobileNumberRegex := regexp.MustCompile(`^(\+628|628|08|8)`)
	if !mobileNumberRegex.MatchString(request.MobileNumber) {
		return helper.CustomError(http.StatusBadRequest, "Wrong format mobile number")
	}
	request.MobileNumber = mobileNumberRegex.ReplaceAllString(request.MobileNumber, "+628")

	// Check username is already registered or not yet
	user, err := service.repository.FindOneUser(ctx, bson.M{
		"username": request.Username,
	})
	if err != nil {
		return helper.CustomError(http.StatusInternalServerError, err.Error())
	}
	if user != nil {
		return helper.CustomError(http.StatusUnprocessableEntity, "Username already registered")
	}

	// Check username is already registered or not yet
	user, err = service.repository.FindOneUser(ctx, bson.M{
		"mobileNumber": request.MobileNumber,
	})
	if err != nil {
		return helper.CustomError(http.StatusInternalServerError, err.Error())
	}
	if user != nil {
		return helper.CustomError(http.StatusUnprocessableEntity, "Mobile number already registered")
	}

	// Generate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	// Generate otp 6 digit
	otp, err := helper.GenerateOTP(6)
	if err != nil {
		return helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	// Generate user json for redis
	userRedis := domain.UserRedis{
		Username:     request.Username,
		Password:     string(hashedPassword),
		Name:         request.Name,
		Otp:          *otp,
		MobileNumber: request.MobileNumber,
	}
	userRedisJson, err := json.Marshal(userRedis)
	if err != nil {
		return helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	// Insert user json to redis
	if err = service.redis.GetClient().Set(
		ctx,
		fmt.Sprintf("REGISTER:%s", request.Username),
		string(userRedisJson),
		time.Minute*10,
	).Err(); err != nil {
		return helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	// Generate user json for kafka
	userKafka := domain.UserKafka{
		Name:         request.Name,
		Otp:          *otp,
		MobileNumber: request.MobileNumber,
	}
	userKafkaJson, err := json.Marshal(userKafka)
	if err != nil {
		return helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	// Produce user json to kafka
	err = service.kafkaProducer.SendMessage("REGISTER-OTP", string(userKafkaJson))
	if err != nil {
		return helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (service ServiceImpl) VerifyRegister(ctx context.Context, request web.VerifyRegisterRequest) (*web.VerifyRegisterResponse, error) {
	request.Username = strings.ToLower(request.Username)

	// Check username is already registered or not yet
	user, err := service.repository.FindOneUser(ctx, bson.M{
		"username": request.Username,
	})
	if err != nil {
		return nil, helper.CustomError(http.StatusInternalServerError, err.Error())
	}
	if user != nil {
		return nil, helper.CustomError(http.StatusUnprocessableEntity, "Username already registered")
	}

	// Get user from redis
	userRedisJson, err := service.redis.GetClient().Get(ctx, fmt.Sprintf("REGISTER:%s", request.Username)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, helper.CustomError(http.StatusUnauthorized, "OTP has expired")
		}
		return nil, helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	// Parse user json to user object
	var userRedis domain.UserRedis
	err = json.Unmarshal([]byte(userRedisJson), &userRedis)
	if err != nil {
		return nil, helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	// Validate otp
	if userRedis.Otp != request.Otp {
		return nil, helper.CustomError(http.StatusUnauthorized, "Invalid OTP")
	}

	// Insert user to mongodb
	insertUser := domain.InsertUser{
		Username:     userRedis.Username,
		Password:     userRedis.Password,
		Name:         userRedis.Name,
		MobileNumber: userRedis.MobileNumber,
	}
	insertedId, err := service.repository.InsertOneUser(ctx, insertUser)
	if err != nil {
		return nil, helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	// Delete user in redis
	err = service.redis.GetClient().Del(ctx, fmt.Sprintf("REGISTER:%s", request.Username)).Err()
	if err != nil {
		return nil, helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	// Generate token
	token, err := service.jwtAuth.GenerateToken(jwt.Payload{
		Id: *insertedId,
	})
	if err != nil {
		return nil, helper.CustomError(http.StatusInternalServerError, err.Error())
	}

	return &web.VerifyRegisterResponse{
		Token: *token,
	}, nil
}
