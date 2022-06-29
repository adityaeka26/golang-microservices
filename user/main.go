package main

import (
	"github.com/adityaeka26/golang-microservices/user/config"
	"github.com/adityaeka26/golang-microservices/user/database"
	"github.com/adityaeka26/golang-microservices/user/jwt"
	"github.com/adityaeka26/golang-microservices/user/kafka"
	"github.com/adityaeka26/golang-microservices/user/middleware"
	"github.com/adityaeka26/golang-microservices/user/module/handler"
	"github.com/adityaeka26/golang-microservices/user/module/repository"
	"github.com/adityaeka26/golang-microservices/user/module/service"
	"github.com/adityaeka26/golang-microservices/user/router"
)

func main() {
	config := config.NewConfig()

	jwtAuth := jwt.NewJWT(config)
	authMiddleware := middleware.NewAuth()

	kafkaProducer := kafka.NewKafkaProducer()
	redis := database.NewRedis(config.GetEnv().RedisUrl, config.GetEnv().RedisPassword)
	mongoDatabase := database.NewMongoDB(config.GetEnv().MongodbUrl, config.GetEnv().MongodbDatabaseName)

	repository := repository.NewRepository(mongoDatabase)
	service := service.NewService(repository, jwtAuth, redis, kafkaProducer)
	handler := handler.NewHandler(service)
	router := router.NewRouter(handler, authMiddleware)

	router.GetGinEngine().Run(":8080")
}
