package products

import (
	"net/http"
	"strconv"

	"ecommerce/database"
	"ecommerce/utils"
)

// Delete Product
func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	p := database.DeleteProduct(id)
	if p != nil {
		utils.HandleSendData(w, database.ProductResponse{
			Message: "Product Deleted successfully!",
			Data:    *p,
		}, http.StatusOK)
		return
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}
