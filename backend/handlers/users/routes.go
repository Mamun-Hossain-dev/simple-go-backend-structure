package users

import (
	"net/http"
)

func (h *Handler) UsersRouter() *http.ServeMux {
	r := http.NewServeMux()

	// Routes
	r.Handle("POST /", http.HandlerFunc(h.CreateUser))
	r.Handle("POST /login", http.HandlerFunc(h.LoginUser))

	return r
}
