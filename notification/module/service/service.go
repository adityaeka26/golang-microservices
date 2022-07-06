package service

import (
	"encoding/json"
	"fmt"

	"github.com/adityaeka26/golang-microservices/notification/logger"
	"github.com/adityaeka26/golang-microservices/notification/module/model/event"
	"github.com/adityaeka26/golang-microservices/notification/module/repository"
	"go.uber.org/zap"
)

type Service interface {
	SendRegisterOtp(payload event.RegisterOtpKafka) error
}

type ServiceImpl struct {
	repository repository.Repository
	logger     logger.Logger
}

func NewService(repository repository.Repository, logger logger.Logger) Service {
	return &ServiceImpl{
		repository: repository,
		logger:     logger,
	}
}

func (service ServiceImpl) SendRegisterOtp(payload event.RegisterOtpKafka) error {
	context := "service-SendRegisterOtp"
	marshaledPayload, err := json.Marshal(payload)
	if err != nil {
		service.logger.GetLogger().Error(
			"Marshal payload fail",
			zap.String("context", context),
			zap.Error(err),
		)
		panic(err)
	}

	fmt.Println(payload)

	service.logger.GetLogger().Info(
		"Send register otp success",
		zap.String("context", context),
		zap.ByteString("payload", marshaledPayload),
	)

	return nil
}
