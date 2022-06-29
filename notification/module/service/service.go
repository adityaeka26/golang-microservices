package service

import "github.com/adityaeka26/golang-microservices/notification/module/repository"

type Service interface{}

type ServiceImpl struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return &ServiceImpl{
		repository: repository,
	}
}
