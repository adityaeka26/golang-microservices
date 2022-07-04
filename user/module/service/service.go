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
	"github.com/adityaeka26/golang-microservices/user/logger"
	"github.com/adityaeka26/golang-microservices/user/module/model/domain"
	"github.com/adityaeka26/golang-microservices/user/module/model/web"
	"github.com/adityaeka26/golang-microservices/user/module/repository"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
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
	logger        logger.Logger
}

func NewService(repository repository.Repository, jwtAuth jwt.JWT, redis database.Redis, kafkaProducer kafka.KafkaProducer, logger logger.Logger) Service {
	return &ServiceImpl{
		repository:    repository,
		jwtAuth:       jwtAuth,
		redis:         redis,
		kafkaProducer: kafkaProducer,
		logger:        logger,
	}
}

func (service ServiceImpl) Register(ctx context.Context, request web.RegisterRequest) error {
	namespace := "service-Register"
	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		service.logger.GetLogger().Error("Marshal request fail", zap.Error(err), zap.Namespace(namespace))
		return helper.CustomError(http.StatusInternalServerError, "Internal server error")
	}

	// Preprocecss username & mobile number
	request.Username = strings.ToLower(request.Username)
	mobileNumberRegex := regexp.MustCompile(`^(\+628|628|08|8)`)
	if !mobileNumberRegex.MatchString(request.MobileNumber) {
		service.logger.GetLogger().Warn("Wrong format mobile number", zap.Namespace(namespace), zap.ByteString("request", marshaledRequest))
		return helper.CustomError(http.StatusBadRequest, "Wrong format mobile number")
	}
	request.MobileNumber = mobileNumberRegex.ReplaceAllString(request.MobileNumber, "+628")

	// Check username is already registered or not yet
	user, err := service.repository.FindOneUser(ctx, bson.M{
		"username": request.Username,
	})
	if err != nil {
		service.logger.GetLogger().Error("Find one user fail", zap.Error(err), zap.Namespace(namespace), zap.ByteString("request", marshaledRequest))
		return helper.CustomError(http.StatusInternalServerError, "Internal server error")
	}
	if user != nil {
		service.logger.GetLogger().Warn("Username already registered", zap.Namespace(namespace), zap.ByteString("request", marshaledRequest))
		return helper.CustomError(http.StatusUnprocessableEntity, "Username already registered")
	}

	// Check username is already registered or not yet
	user, err = service.repository.FindOneUser(ctx, bson.M{
		"mobileNumber": request.MobileNumber,
	})
	if err != nil {
		service.logger.GetLogger().Error("Find one user 2 fail", zap.Error(err), zap.Namespace(namespace), zap.ByteString("request", marshaledRequest))
		return helper.CustomError(http.StatusInternalServerError, "Internal server error")
	}
	if user != nil {
		service.logger.GetLogger().Warn("Mobile number already registered", zap.Namespace(namespace), zap.ByteString("request", marshaledRequest))
		return helper.CustomError(http.StatusUnprocessableEntity, "Mobile number already registered")
	}

	// Generate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		service.logger.GetLogger().Error("Generate hashed password fail", zap.Error(err), zap.Namespace(namespace), zap.ByteString("request", marshaledRequest))
		return helper.CustomError(http.StatusInternalServerError, "Internal server error")
	}

	// Generate otp 6 digit
	otp, err := helper.GenerateOTP(6)
	if err != nil {
		service.logger.GetLogger().Error("Generate otp fail", zap.Error(err), zap.Namespace(namespace), zap.ByteString("request", marshaledRequest))
		return helper.CustomError(http.StatusInternalServerError, "Internal server error")
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
		service.logger.GetLogger().Error("Marshal redis fail", zap.Error(err), zap.Namespace(namespace), zap.ByteString("request", marshaledRequest))
		return helper.CustomError(http.StatusInternalServerError, "Internal server error")
	}

	// Insert user json to redis
	if err = service.redis.GetClient().Set(
		ctx,
		fmt.Sprintf("REGISTER:%s", request.Username),
		string(userRedisJson),
		time.Minute*10,
	).Err(); err != nil {
		service.logger.GetLogger().Error("Set redis fail", zap.Error(err), zap.Namespace(namespace), zap.ByteString("request", marshaledRequest))
		return helper.CustomError(http.StatusInternalServerError, "Internal server error")
	}

	// Generate user json for kafka
	userKafka := domain.RegisterOtpKafka{
		Name:         request.Name,
		Otp:          *otp,
		MobileNumber: request.MobileNumber,
	}
	userKafkaJson, err := json.Marshal(userKafka)
	if err != nil {
		service.logger.GetLogger().Error("Marshal kafka fail", zap.Error(err), zap.Namespace(namespace), zap.ByteString("request", marshaledRequest))
		return helper.CustomError(http.StatusInternalServerError, "Internal server error")
	}

	// Produce user json to kafka
	err = service.kafkaProducer.SendMessage("REGISTER-OTP", string(userKafkaJson))
	if err != nil {
		service.logger.GetLogger().Error("Send message kafka fail", zap.Error(err), zap.Namespace(namespace), zap.ByteString("request", marshaledRequest))
		return helper.CustomError(http.StatusInternalServerError, "Internal server error")
	}

	service.logger.GetLogger().Info(
		"Register success",
		zap.Namespace(namespace),
		zap.ByteString("request", marshaledRequest),
	)

	return nil
}

func (service ServiceImpl) VerifyRegister(ctx context.Context, request web.VerifyRegisterRequest) (*web.VerifyRegisterResponse, error) {
	namespace := "service-VerifyRegister"
	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		service.logger.GetLogger().Error(
			"Marshal request fail",
			zap.Namespace(namespace),
			zap.Error(err),
		)
		return nil, helper.CustomError(http.StatusInternalServerError, "Internal server error")
	}

	request.Username = strings.ToLower(request.Username)

	// Check username is already registered or not yet
	user, err := service.repository.FindOneUser(ctx, bson.M{
		"username": request.Username,
	})
	if err != nil {
		service.logger.GetLogger().Error(
			"Find one user fail",
			zap.Namespace(namespace),
			zap.Error(err),
			zap.ByteString("request", marshaledRequest),
		)
		return nil, helper.CustomError(http.StatusInternalServerError, "Internal server error")
	}
	if user != nil {
		service.logger.GetLogger().Warn(
			"Username already registered",
			zap.Namespace(namespace),
			zap.ByteString("request", marshaledRequest),
		)
		return nil, helper.CustomError(http.StatusUnprocessableEntity, "Username already registered")
	}

	// Get user from redis
	userRedisJson, err := service.redis.GetClient().Get(ctx, fmt.Sprintf("REGISTER:%s", request.Username)).Result()
	if err != nil {
		if err == redis.Nil {
			service.logger.GetLogger().Warn(
				"OTP has expired",
				zap.Namespace(namespace),
				zap.ByteString("request", marshaledRequest),
			)
			return nil, helper.CustomError(http.StatusUnauthorized, "OTP has expired")
		}
		service.logger.GetLogger().Error(
			"Get redis fail",
			zap.Namespace(namespace),
			zap.Error(err),
			zap.ByteString("request", marshaledRequest),
		)
		return nil, helper.CustomError(http.StatusInternalServerError, "Internal server error")
	}

	// Parse user json to user object
	var userRedis domain.UserRedis
	err = json.Unmarshal([]byte(userRedisJson), &userRedis)
	if err != nil {
		service.logger.GetLogger().Error(
			"Unmarshal redis fail",
			zap.Namespace(namespace),
			zap.Error(err),
			zap.ByteString("request", marshaledRequest),
		)
		return nil, helper.CustomError(http.StatusInternalServerError, "Internal server error")
	}

	// Validate otp
	if userRedis.Otp != request.Otp {
		service.logger.GetLogger().Warn(
			"Invalid OTP",
			zap.Namespace(namespace),
			zap.ByteString("request", marshaledRequest),
		)
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
		service.logger.GetLogger().Error(
			"Insert one user fail",
			zap.Namespace(namespace),
			zap.Error(err),
			zap.ByteString("request", marshaledRequest),
		)
		return nil, helper.CustomError(http.StatusInternalServerError, "Internal server error")
	}

	// Delete user in redis
	err = service.redis.GetClient().Del(ctx, fmt.Sprintf("REGISTER:%s", request.Username)).Err()
	if err != nil {
		service.logger.GetLogger().Error(
			"Delete redis fail",
			zap.Namespace(namespace),
			zap.Error(err),
			zap.ByteString("request", marshaledRequest),
		)
		return nil, helper.CustomError(http.StatusInternalServerError, "Internal server error")
	}

	// Generate token
	token, err := service.jwtAuth.GenerateToken(jwt.Payload{
		Id: *insertedId,
	})
	if err != nil {
		service.logger.GetLogger().Error(
			"Generate token fail",
			zap.Namespace(namespace),
			zap.Error(err),
			zap.ByteString("request", marshaledRequest),
		)
		return nil, helper.CustomError(http.StatusInternalServerError, "Internal server error")
	}

	response := web.VerifyRegisterResponse{
		Token: *token,
	}
	marshaledResponse, err := json.Marshal(response)
	if err != nil {
		service.logger.GetLogger().Error(
			"Marshal response fail",
			zap.Namespace(namespace),
			zap.Error(err),
			zap.ByteString("request", marshaledRequest),
		)
		return nil, helper.CustomError(http.StatusInternalServerError, "Internal server error")
	}
	service.logger.GetLogger().Info(
		"Verify register success",
		zap.Namespace(namespace),
		zap.ByteString("request", marshaledRequest),
		zap.ByteString("response", marshaledResponse),
	)

	return &response, nil
}
