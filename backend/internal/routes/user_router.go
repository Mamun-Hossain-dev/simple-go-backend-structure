package routes

import (
	"net/http"

	"ecommerce/internal/user"
)

func UsersRouter() *http.ServeMux {
	r := http.NewServeMux()

	repo := user.NewUserRepository()
	service := user.NewUserService(repo)
	h := user.NewHandler(service)

	r.Handle("POST /", http.HandlerFunc(h.Register))
	r.Handle("POST /login", http.HandlerFunc(h.Login))

	return r
}
