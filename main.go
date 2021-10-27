package main

import (
	"fmt"
	"net/http"
	"rest-poc/controller"
	"rest-poc/router"
	"rest-poc/service"
)

var (
	svc               service.ProductService       = service.NewProductService()
	productController controller.ProductController = controller.NewProductController(svc)
	httpRoute         router.Router                = router.NewMuxRouter()
)

func main() {

	const port string = ":8000"
	httpRoute.GET("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Ok!!")
	})

	httpRoute.GET("/products", productController.GeAllProducts)
	httpRoute.POST("/products", productController.AddProduct)
	httpRoute.PUT("/products", productController.UpdateProduct)
	httpRoute.DELETE("/products/{id}", productController.Delete)

	httpRoute.SERVE(port)
}
