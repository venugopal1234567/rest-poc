package repository

import "rest-poc/entity"

type ProductRpository interface {
	List() ([]*entity.Product, error)
	Add(p *entity.Product) error
}
