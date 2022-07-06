package main

import (
	"github.com/adityaeka26/golang-microservices/notification/config"
	"github.com/adityaeka26/golang-microservices/notification/kafka"
	"github.com/adityaeka26/golang-microservices/notification/logger"
	"github.com/adityaeka26/golang-microservices/notification/module/handler"
	"github.com/adityaeka26/golang-microservices/notification/module/repository"
	"github.com/adityaeka26/golang-microservices/notification/module/service"
	"github.com/adityaeka26/golang-microservices/notification/router"
)

func main() {
	config := config.NewConfig()

	logger := logger.NewLog(config.GetEnv().AppName, config.GetEnv().AppEnvironment)

	kafkaConsumer := kafka.NewKafkaConsumer(config.GetEnv().KafkaUrl, logger)

	repository := repository.NewRepository()
	service := service.NewService(repository, logger)
	handler := handler.NewHandler(service, kafkaConsumer, logger)
	router := router.NewRouter(handler)

	router.GetGinEngine().Run(":8082")
}
