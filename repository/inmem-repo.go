package repository

import "rest-poc/entity"

type repo struct{}

//NewInMemRepository
func NewInMemRepository() ProductRpository {
	return &repo{}
}

var (
	products []*entity.Product
)

func init() {
	products = []*entity.Product{{
		ID:       "12add",
		Name:     "samsung",
		Cost:     300,
		Quantity: 12,
	}}
}

func (ps *repo) List() ([]*entity.Product, error) {
	return products, nil
}

func (ps *repo) Add(p *entity.Product) error {
	products = append(products, p)
	return nil
}
