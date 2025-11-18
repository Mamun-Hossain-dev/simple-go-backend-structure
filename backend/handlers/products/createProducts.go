package products

import (
	"encoding/json"
	"net/http"

	"ecommerce/database"
	"ecommerce/utils"
)

// POST /create-product
func (h *Handler) CreateProducts(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	createdProduct := database.StoreProduct(newProduct)
	response := database.ProductResponse{
		Message: "Product created successfully!",
		Data:    createdProduct,
	}

	utils.HandleSendData(w, response, http.StatusCreated)
}
