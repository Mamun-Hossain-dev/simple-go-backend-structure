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

	mainMux.Handle("/api/users/",
		utils.ChainMiddlewares(
			// FIX: strip "/api/users/" OR "/api/users"
			http.StripPrefix("/api/users", userRouter),
			middlewares.CorsMiddleware,
			middlewares.LoggingMiddleware,
		),
	)

	// ----------------------------
	// PRODUCTS ROUTER
	// ----------------------------
	productRouter := routes.ProductRouter(db)

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
