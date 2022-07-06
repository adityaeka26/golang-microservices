package handler

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/adityaeka26/golang-microservices/notification/kafka"
	"github.com/adityaeka26/golang-microservices/notification/logger"
	"github.com/adityaeka26/golang-microservices/notification/module/model/event"
	"github.com/adityaeka26/golang-microservices/notification/module/service"
	"go.uber.org/zap"
)

type Handler interface {
	SendRegisterOtp()
}

type HandlerImpl struct {
	service       service.Service
	kafkaConsumer kafka.KafkaConsumer
	logger        logger.Logger
}

func NewHandler(service service.Service, kafkaConsumer kafka.KafkaConsumer, logger logger.Logger) Handler {
	return &HandlerImpl{
		service:       service,
		kafkaConsumer: kafkaConsumer,
		logger:        logger,
	}
}

func (handler HandlerImpl) SendRegisterOtp() {
	context := "handler-SendRegisterOtp"

	chanMessage := make(chan sarama.ConsumerMessage, 256)
	go handler.kafkaConsumer.Consume("REGISTER-OTP", chanMessage)
	for msg := range chanMessage {
		payload := event.RegisterOtpKafka{}
		if err := json.Unmarshal(msg.Value, &payload); err != nil {
			handler.logger.GetLogger().Error(
				"Unmarshal data from kafka error",
				zap.String("context", context),
				zap.Error(err),
			)
			panic(err)
		}
		if err := handler.service.SendRegisterOtp(payload); err != nil {
			panic(err)
		}
	}
}
