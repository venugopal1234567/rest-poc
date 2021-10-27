package service

import (
	"rest-poc/entity"
	"rest-poc/repository"
)

type ProductService interface {
	Add(product *entity.Product) error
	List() ([]*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}

type service struct{}

func NewProductService() ProductService {
	return &service{}
}

var (
	repo repository.ProductRpository = repository.NewMongoRepository()
)

func (s *service) Add(product *entity.Product) error {
	return repo.Add(product)
}

func (s *service) List() ([]*entity.Product, error) {
	return repo.List()
}

func (s *service) Update(product *entity.Product) error {
	return repo.Update(product)
}

func (s *service) Delete(id string) error {
	return repo.Delete(id)
}
