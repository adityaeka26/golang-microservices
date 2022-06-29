package handler

import "github.com/adityaeka26/golang-microservices/notification/module/service"

type Handler interface{}

type HandlerImpl struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return &HandlerImpl{
		service: service,
	}
}
