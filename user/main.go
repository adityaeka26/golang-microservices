package main

import (
	"github.com/adityaeka26/golang-microservices/user/config"
	"github.com/adityaeka26/golang-microservices/user/database"
	"github.com/adityaeka26/golang-microservices/user/jwt"
	"github.com/adityaeka26/golang-microservices/user/module/handler"
	"github.com/adityaeka26/golang-microservices/user/module/repository"
	"github.com/adityaeka26/golang-microservices/user/module/service"
	"github.com/adityaeka26/golang-microservices/user/router"
)

func main() {
	config := config.NewConfig()
	jwtAuth := jwt.NewJWT()
	mongoDatabase := database.NewMongoDB(config.GetEnv().MongodbUrl, config.GetEnv().MongodbDatabaseName)
	repository := repository.NewRepository(mongoDatabase)
	service := service.NewService(repository, jwtAuth)
	handler := handler.NewHandler(service)
	router := router.NewRouter(handler)

	router.GetGinEngine().Run(":8080")
}
