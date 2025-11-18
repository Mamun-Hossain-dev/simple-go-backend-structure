package products

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ecommerce/database"
	"ecommerce/utils"
)

// Update Product
func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	StrId := r.PathValue("id")
	id, err := strconv.Atoi(StrId)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct database.Product

	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	p := database.UpdateProduct(id, updatedProduct)
	if p != nil {
		utils.HandleSendData(w, database.ProductResponse{
			Message: "Product Updated successfully!",
			Data:    *p,
		}, http.StatusOK)
		return
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}
