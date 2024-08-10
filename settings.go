package amunet_auth_microservices

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBName     string
	DBPort     string
	DBUsername string
	DBPassword string
	DBHost     string
	SSLMode    string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic("failed to load env variables")
	}
	return &Config{
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		SSLMode:    os.Getenv("SSL_MODE"),
	}
}
