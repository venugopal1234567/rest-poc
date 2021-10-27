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
		Name:     "samsung",
		Cost:     300,
		Quantity: 12,
	}}, nil
}

func (s *ServiceMock) Update(product *entity.Product) error {
	if s.ErrResp != nil {
		return s.ErrResp
	}
	return nil
}

func (s *ServiceMock) Delete(id string) error {
	if s.ErrResp != nil {
		return s.ErrResp
	}
	return nil
}
