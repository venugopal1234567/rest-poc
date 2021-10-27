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

func (ps *repo) Update(p *entity.Product) error {
	return nil
}

func (ps *repo) Delete(id string) error {
	return nil
}
