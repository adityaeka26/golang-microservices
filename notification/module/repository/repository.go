package repository

type Repository interface{}

type RepositoryImpl struct{}

func NewRepository() Repository {
	return &RepositoryImpl{}
}
