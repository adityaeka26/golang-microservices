package main

import (
	"github.com/adityaeka26/golang-microservices/user/config"
	"github.com/adityaeka26/golang-microservices/user/database"
	"github.com/adityaeka26/golang-microservices/user/jwt"
	"github.com/adityaeka26/golang-microservices/user/kafka"
	"github.com/adityaeka26/golang-microservices/user/logger"
	"github.com/adityaeka26/golang-microservices/user/middleware"
	"github.com/adityaeka26/golang-microservices/user/module/handler"
	"github.com/adityaeka26/golang-microservices/user/module/repository"
	"github.com/adityaeka26/golang-microservices/user/module/service"
	"github.com/adityaeka26/golang-microservices/user/router"
)

func main() {
	config := config.NewConfig()

	// log.Add

	jwtAuth := jwt.NewJWT(config)
	authMiddleware := middleware.NewAuth()
	logger := logger.NewLog(config.GetEnv().AppName, config.GetEnv().AppEnvironment)

	kafkaProducer := kafka.NewKafkaProducer(config.GetEnv().KafkaUrl, logger)
	redis := database.NewRedis(config.GetEnv().RedisUrl, config.GetEnv().RedisPassword)
	mongoDatabase := database.NewMongoDB(config.GetEnv().MongodbUrl, config.GetEnv().MongodbDatabaseName)

	repository := repository.NewRepository(mongoDatabase)
	service := service.NewService(repository, jwtAuth, redis, kafkaProducer, logger)
	handler := handler.NewHandler(service, logger)
	router := router.NewRouter(handler, authMiddleware)

	router.GetGinEngine().Run(":8080")
}
