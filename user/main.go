package main

import (
	"github.com/adityaeka26/golang-microservices/user/app"
	"github.com/adityaeka26/golang-microservices/user/module/query/handler"
)

func main() {
	_ = app.NewMongoDB("mongodb://root:leomessi@localhost:27017/book_db?authSource=admin&ssl=false")
	handler := handler.NewHandler()
	router := app.NewRouter(handler)

	router.Run(":8080")
}
