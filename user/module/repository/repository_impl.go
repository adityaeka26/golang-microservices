package repository

type RepositoryImpl struct{}

func NewRepository() Repository {
	return &RepositoryImpl{}
}
