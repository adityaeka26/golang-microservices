package handler

import (
	"encoding/json"
	"os"

	"github.com/Shopify/sarama"
	"github.com/adityaeka26/golang-microservices/notification/kafka"
	"github.com/adityaeka26/golang-microservices/notification/module/model/event"
	"github.com/adityaeka26/golang-microservices/notification/module/service"
)

type Handler interface {
	SendRegisterOtp()
}

type HandlerImpl struct {
	service       service.Service
	kafkaConsumer kafka.KafkaConsumer
}

func NewHandler(service service.Service, kafkaConsumer kafka.KafkaConsumer) Handler {
	return &HandlerImpl{
		service:       service,
		kafkaConsumer: kafkaConsumer,
	}
}

func (handler HandlerImpl) SendRegisterOtp() {
	signals := make(chan os.Signal, 1)
	chanMessage := make(chan *sarama.ConsumerMessage, 256)
	go handler.kafkaConsumer.Consume("REGISTER-OTP", chanMessage)

loop:
	for {
		select {
		case msg := <-chanMessage:
			payload := event.RegisterOtpKafka{}
			if err := json.Unmarshal(msg.Value, &payload); err != nil {
				panic(err)
			}
			if err := handler.service.SendRegisterOtp(payload); err != nil {
				panic(err)
			}
		case sig := <-signals:
			if sig == os.Interrupt {
				break loop
			}
		}
	}
}
