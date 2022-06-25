package main

import (
	"github.com/adityaeka26/golang-microservices/user/app"
	"github.com/adityaeka26/golang-microservices/user/module/handler"
	"github.com/adityaeka26/golang-microservices/user/module/service"
)

func main() {
	_ = app.NewMongoDB("mongodb://root:leomessi@localhost:27017/book_db?authSource=admin&ssl=false")
	service := service.NewService()
	handler := handler.NewHandler(service)
	router := app.NewRouter(handler)

	router.Run(":8080")
}
