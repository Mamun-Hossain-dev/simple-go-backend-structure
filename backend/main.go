package main

import (
	"fmt"
	"log"
	"net/http"

	"ecommerce/internal/config"
	"ecommerce/internal/product"
	"ecommerce/internal/routes"
	"ecommerce/middlewares"
	"ecommerce/utils"
)

func main() {
	// Load config from .env
	cfg := config.LoadConfig()

	// Create main mux
	mainMux := http.NewServeMux()

	// Load fake product database
	product.LoadFakeProducts()

	// ----------------------------
	// USERS ROUTER
	// ----------------------------
	userRouter := routes.UsersRouter()

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
	productRouter := routes.ProductRouter()

	mainMux.Handle("/api/products/",
		utils.ChainMiddlewares(
			http.StripPrefix("/api/products", productRouter),
			middlewares.CorsMiddleware,
			middlewares.LoggingMiddleware,
			// middlewares.AuthMiddleware, // uncomment if you want JWT protected
		),
	)

	// ----------------------------
	// START SERVER
	// ----------------------------
	fmt.Println("🚀 Server running on port:", cfg.HttpPort)

	log.Fatal(http.ListenAndServe(":"+cfg.HttpPort, mainMux))
}
