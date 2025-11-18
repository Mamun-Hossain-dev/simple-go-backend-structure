package main

import (
	"ecommerce/database"
	"ecommerce/server"
)

func main() {
	// Start The Server
	server.Server()
}

// int function
func init() {
	database.LoadProducts()
}
