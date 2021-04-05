package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	KeycloakConfig KeycloakConfig
	HTTP           HTTPConfig
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		KeycloakConfig: LoadKeycloakConfig(),
		HTTP:           LoadHTTPConfig(),
	}
}
