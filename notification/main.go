package main

import (
	"github.com/adityaeka26/golang-microservices/notification/config"
	"github.com/adityaeka26/golang-microservices/notification/kafka"
	"github.com/adityaeka26/golang-microservices/notification/module/handler"
	"github.com/adityaeka26/golang-microservices/notification/module/repository"
	"github.com/adityaeka26/golang-microservices/notification/module/service"
	"github.com/adityaeka26/golang-microservices/notification/router"
)

func main() {
	_ = config.NewConfig()

	kafkaConsumer := kafka.NewKafkaConsumer()

	repository := repository.NewRepository()
	service := service.NewService(repository)
	handler := handler.NewHandler(service, kafkaConsumer)

	router := router.NewRouter(handler)

	router.GetGinEngine().Run(":8082")
}
