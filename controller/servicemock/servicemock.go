package servicemock

import (
	"rest-poc/entity"
)

type ServiceMock struct {
	ErrResp error
}

func (s *ServiceMock) Add(product *entity.Product) error {
	if s.ErrResp != nil {
		return s.ErrResp
	}
	return nil
}

func (s *ServiceMock) List() ([]*entity.Product, error) {
	if s.ErrResp != nil {
		return []*entity.Product{}, s.ErrResp
	}
	return []*entity.Product{{
		ID:       "12add",
		Name:     "samsung",
		Cost:     300,
		Quantity: 12,
	}}, nil
}
