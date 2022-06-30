package service

import (
	"fmt"

	"github.com/adityaeka26/golang-microservices/notification/module/model/event"
	"github.com/adityaeka26/golang-microservices/notification/module/repository"
)

type Service interface {
	SendRegisterOtp(payload event.RegisterOtpKafka) error
}

type ServiceImpl struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return &ServiceImpl{
		repository: repository,
	}
}

func (service ServiceImpl) SendRegisterOtp(payload event.RegisterOtpKafka) error {
	fmt.Println(payload)
	return nil
}
