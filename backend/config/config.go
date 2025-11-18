package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    string
	Jwt_secret  string
}

var configurations *Config

func LoadConfig() *Config {
	// if config already loaded return it
	if configurations != nil {
		return configurations
	}

	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading from system environment")
	}

	version := os.Getenv("VERSION")
	if version == "" {
		log.Println("VERSION env is required!")
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Println("SERVICE_NAME env is required!")
	}

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		log.Println("PORT env is required!")
	}

	jwt_secret := os.Getenv("JWT_SECRET")
	if jwt_secret == "" {
		log.Println("JWT_SECRET env is required!")
	}

	configurations = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    port,
		Jwt_secret:  jwt_secret,
	}

	return configurations
}
