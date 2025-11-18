package products

import (
	"net/http"

	"ecommerce/middlewares"
)

func (h *Handler) ProductRouter() *http.ServeMux {
	r := http.NewServeMux()

	// Public Routes
	r.Handle("GET /", http.HandlerFunc(h.GetProducts))
	r.Handle("GET /{id}", http.HandlerFunc(h.GetSingleProduct))

	// Protected Routes
	r.Handle("POST /", middlewares.AuthMiddleware(http.HandlerFunc(h.CreateProducts)))
	r.Handle("PUT /{id}", middlewares.AuthMiddleware(http.HandlerFunc(h.UpdateProduct)))
	r.Handle("DELETE /{id}", middlewares.AuthMiddleware(http.HandlerFunc(h.DeleteProduct)))

	return r
}
