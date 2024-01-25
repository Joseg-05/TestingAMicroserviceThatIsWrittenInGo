package config

import (
	"log"
	"os"

	"authservice/internal/helper"

	"github.com/joho/godotenv"
)

func SetupEnvironment() string {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	//setup jwt helper
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		log.Fatal("secret is not specified")
	}

	helper.InitJwtHelper(secret)

	//setup and return port
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port is not specified")
	}
	return portString
}
