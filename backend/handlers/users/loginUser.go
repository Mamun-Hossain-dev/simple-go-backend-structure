package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/utils"
)

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loggedUser database.LoggedUser

	// decode JSON body
	err := json.NewDecoder(r.Body).Decode(&loggedUser)
	if err != nil {
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	// find user
	user := database.Find(loggedUser.Email, loggedUser.Password)
	if user == nil {
		http.Error(w, "Invalid Credential", http.StatusUnauthorized)
		return
	}

	secret := []byte(config.LoadConfig().Jwt_secret)

	// FIX: convert int → string properly
	userID := strconv.Itoa(user.ID)

	// generate token
	token, err := utils.CreateToken(secret, userID, "admin", user.FirstName)
	if err != nil {
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	// response
	response := map[string]interface{}{
		"message":     "Login successful",
		"user":        user,
		"accessToken": token,
	}

	utils.HandleSendData(w, response, http.StatusOK)
}
