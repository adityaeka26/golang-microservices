package service

type ServiceImpl struct{}

func NewService() Service {
	return &ServiceImpl{}
}
