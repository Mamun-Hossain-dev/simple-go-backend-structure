package server

import (
	"fmt"
	"log"
	"net/http"

	"ecommerce/config"
	"ecommerce/handlers/products"
	"ecommerce/handlers/users"
	"ecommerce/middlewares"

	"ecommerce/utils"
)

func Server() {
	mux := http.NewServeMux()

	// Load config
	cnf := config.LoadConfig()

	// subRouters
	productRouter := products.NewHandler().ProductRouter()
	userRouter := users.NewHandler().UsersRouter()

	// Apply middleware chain per group
	mux.Handle("/api/products/", utils.ChainMiddlewares(
		http.StripPrefix("/api/products", productRouter),
		middlewares.CorsMiddleware,
		middlewares.LoggingMiddleware,
	))

	mux.Handle("/api/users/", utils.ChainMiddlewares(
		http.StripPrefix("/api/users", userRouter),
		middlewares.CorsMiddleware,
		middlewares.LoggingMiddleware,
	))

	fmt.Println("Server is running on port: ", cnf.HttpPort)

	if err := http.ListenAndServe(":"+cnf.HttpPort, mux); err != nil {
		log.Fatal("Failed to Start The Server", err)
	}
}
