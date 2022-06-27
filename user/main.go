package main

import (
	"github.com/adityaeka26/golang-microservices/user/app"
	"github.com/adityaeka26/golang-microservices/user/database"
	"github.com/adityaeka26/golang-microservices/user/module/handler"
	"github.com/adityaeka26/golang-microservices/user/module/repository"
	"github.com/adityaeka26/golang-microservices/user/module/service"
)

func main() {
	mongoDatabase := database.NewMongoDB("mongodb://root:leomessi@localhost:27017/book_db?authSource=admin&ssl=false", "book_db")
	repository := repository.NewRepository(mongoDatabase)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)
	router := app.NewRouter(handler)

	router.Run(":8080")
}
