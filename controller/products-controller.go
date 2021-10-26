package controller

import (
	"encoding/json"
	"net/http"
	"rest-poc/entity"
	"rest-poc/service"
)

type ProductController interface {
	GeAllProducts(rw http.ResponseWriter, r *http.Request)
	AddProduct(rw http.ResponseWriter, r *http.Request)
}

type controller struct {
	svc service.ProductService
}

func NewProductController(svc service.ProductService) ProductController {
	return &controller{
		svc: svc,
	}
}

func (c controller) GeAllProducts(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	products, err := c.svc.List()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{ "error" : "Error retreiving products"}`))
		return
	}

	result, _ := json.Marshal(products)
	rw.WriteHeader(http.StatusOK)
	rw.Write(result)
}

func (c controller) AddProduct(rw http.ResponseWriter, r *http.Request) {
	var product *entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`{ "error" : "Error un marshaling request"}`))
		return
	}
	err = c.svc.Add(product)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{ "error" : "failed to save product"}`))
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{ "message" : "Successfully added product"}`))
}
