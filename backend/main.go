package main

import (
	"fmt"
	"log"
	"net/http"

	"ecommerce/infra/db"
	"ecommerce/internal/config"
	"ecommerce/internal/routes"
	"ecommerce/middlewares"
	"ecommerce/utils"
)

func main() {
	// Load config from .env
	cfg := config.LoadConfig()

	// Initialize DB connection
	db, err := db.NewConnection()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("DB Connection established:", db != nil)

	// Create main mux
	mainMux := http.NewServeMux()

	// ----------------------------
	// USERS ROUTER
	// ----------------------------
	userRouter := routes.UsersRouter(db, cfg)

	// Redirect /api/users → /api/users/
	mainMux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api/users/", http.StatusTemporaryRedirect)
	})

	// Single handler for all subroutes
	mainMux.Handle("/api/users/",
		utils.ChainMiddlewares(
			http.StripPrefix("/api/users", userRouter),
			middlewares.CorsMiddleware,
			middlewares.LoggingMiddleware,
		),
	)

	// ----------------------------
	// PRODUCTS ROUTER
	// ----------------------------
	productRouter := routes.ProductRouter(db)

	// redirect router
	mainMux.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api/products/", http.StatusTemporaryRedirect)
	})

	mainMux.Handle("/api/products/",
		utils.ChainMiddlewares(
			// FIX: REMOVE trailing slash from prefix
			http.StripPrefix("/api/products", productRouter),
			middlewares.CorsMiddleware,
			middlewares.LoggingMiddleware,
		),
	)

	// ----------------------------
	// START SERVER
	// ----------------------------
	fmt.Println("🚀 Server running on port:", cfg.HttpPort)

	log.Fatal(http.ListenAndServe(":"+cfg.HttpPort, mainMux))
}
