package products

import (
	"net/http"

	"ecommerce/database"
	"ecommerce/utils"
)

// GET /products
func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList := database.GetProductList()
	response := database.ProductsResponse{
		Message: "Products data fetch successfully!",
		Data:    productList,
	}
	utils.HandleSendData(w, response, http.StatusOK)
}
