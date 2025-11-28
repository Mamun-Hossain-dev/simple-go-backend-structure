package routes

import (
	"net/http"

	"ecommerce/internal/config"
	"ecommerce/internal/user"

	"github.com/jmoiron/sqlx"
)

func UsersRouter(db *sqlx.DB, cfg *config.Config) *http.ServeMux {
	r := http.NewServeMux()

	repo := user.NewUserRepository(db)
	service := user.NewUserService(repo)
	h := user.NewHandler(service, cfg)

	r.Handle("POST /", http.HandlerFunc(h.Register))
	r.Handle("POST /login", http.HandlerFunc(h.Login))

	return r
}
