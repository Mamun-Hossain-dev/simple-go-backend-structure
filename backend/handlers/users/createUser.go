package users

import (
	"encoding/json"
	"net/http"

	"ecommerce/database"
	"ecommerce/utils"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	createdUser := newUser.StoreUser()
	response := database.UserResponse{
		Message: "user created successfully",
		Data:    createdUser,
	}
	utils.HandleSendData(w, response, 201)
}
