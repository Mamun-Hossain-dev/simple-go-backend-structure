package routes

import (
	"net/http"

	"ecommerce/internal/product"
	"ecommerce/middlewares"
)

func ProductRouter() *http.ServeMux {
	router := http.NewServeMux()

	repo := product.NewProductRepository()
	service := product.NewProductService(repo)
	handler := product.NewProductHandler(service)

	// PUBLIC ROUTES
	router.HandleFunc("GET /", handler.GetProducts)
	router.HandleFunc("GET /{id}", handler.GetProduct)

	// PROTECTED ROUTES (POST, PUT, DELETE)
	router.Handle("POST /",
		middlewares.AuthMiddleware(http.HandlerFunc(handler.CreateProduct)),
	)

	router.Handle("PUT /{id}",
		middlewares.AuthMiddleware(http.HandlerFunc(handler.UpdateProduct)),
	)

	router.Handle("DELETE /{id}",
		middlewares.AuthMiddleware(http.HandlerFunc(handler.DeleteProduct)),
	)

	return router
}
