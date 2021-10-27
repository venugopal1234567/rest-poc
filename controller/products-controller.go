package controller

import (
	"encoding/json"
	"net/http"
	"rest-poc/entity"
	"rest-poc/service"

	"github.com/gorilla/mux"
)

type ProductController interface {
	GeAllProducts(rw http.ResponseWriter, r *http.Request)
	AddProduct(rw http.ResponseWriter, r *http.Request)
	UpdateProduct(rw http.ResponseWriter, r *http.Request)
	Delete(rw http.ResponseWriter, r *http.Request)
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

func (c controller) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	var product *entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`{ "error" : "Error un marshaling request"}`))
		return
	}
	err = c.svc.Update(product)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{ "error" : "failed to update product"}`))
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{ "message" : "Successfully updated product"}`))
}

func (c controller) Delete(rw http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	err := c.svc.Delete(id)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{ "error" : "failed to delete product"}`))
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{ "message" : "Successfully deleted product"}`))
}
