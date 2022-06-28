package service

import (
	"context"

	"github.com/adityaeka26/golang-microservices/user/module/model/web"
)

type Service interface {
	Register(ctx context.Context, request web.RegisterRequest) error
	VerifyRegister(ctx context.Context, request web.VerifyRegisterRequest) (*web.VerifyRegisterResponse, error)
}
