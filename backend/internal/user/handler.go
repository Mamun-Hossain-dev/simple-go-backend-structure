package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ecommerce/internal/config"
	"ecommerce/utils"
)

type Handler struct {
	service UserService
}

func NewHandler(s UserService) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var cu CreateUser
	if err := json.NewDecoder(r.Body).Decode(&cu); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdUser, err := h.service.RegisterUser(cu)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	response := UserResponse{
		Message: "User registered successfully",
		Data:    createdUser,
	}

	utils.HandleSendData(w, response, http.StatusCreated)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var lu LoggedUser
	if err := json.NewDecoder(r.Body).Decode(&lu); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := h.service.LoginUser(lu)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	cfg := config.LoadConfig()
	secret := []byte(cfg.Jwt_secret)

	token, err := utils.CreateToken(secret, strconv.Itoa(user.ID), "user", user.FirstName)
	if err != nil {
		http.Error(w, "token error", 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":     "Login success",
		"user":        user,
		"accessToken": token,
	})
}
