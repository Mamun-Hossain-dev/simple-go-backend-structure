package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ecommerce/utils"
)

type ProductHandler struct {
	service ProductService
}

func NewProductHandler(service ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// All Products
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	data := h.service.GetProducts()
	response := Response{
		Message: "Products retrieved successfully",
		Data:    data,
	}
	utils.HandleSendData(w, response, http.StatusOK)
}

// Single Product by ID
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	p, err := h.service.GetProductByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	response := ProductResponse{
		Message: "Product retrieved successfully",
		Data:    *p,
	}
	utils.HandleSendData(w, response, http.StatusOK)
}

// Create Product
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdProduct := h.service.CreateProduct(newProduct)
	response := ProductResponse{
		Message: "Product created successfully",
		Data:    createdProduct,
	}
	utils.HandleSendData(w, response, http.StatusCreated)
}

// Update Product
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var updatedProduct Product
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	updated, err := h.service.UpdateProduct(id, updatedProduct)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	response := ProductResponse{
		Message: "Product updated successfully",
		Data:    *updated,
	}
	utils.HandleSendData(w, response, http.StatusOK)
}

// Delete Product
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	deleted, err := h.service.DeleteProduct(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	response := ProductResponse{
		Message: "Product deleted successfully",
		Data:    *deleted,
	}
	utils.HandleSendData(w, response, http.StatusOK)
}
