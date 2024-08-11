package amunet_auth_microservices

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	Env           string
	DBName        string
	DBPort        string
	DBUsername    string
	DBPassword    string
	DBHost        string
	SSLMode       string
	AccessTTL     int64
	RefreshTTL    int64
	AccessSecret  string
	RefreshSecret string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic("failed to load env variables")
	}
	accessTTL, err := strconv.ParseInt(os.Getenv("ACCESS_TTL"), 10, 64)
	if err != nil {
		panic("Failed to parse accessTTL")
	}

	refreshTTL, err := strconv.ParseInt(os.Getenv("REFRESH_TTL"), 10, 64)
	if err != nil {
		panic("Failed to parse refreshTTL")
	}
	return &Config{
		Env:           os.Getenv("ENV"),
		DBName:        os.Getenv("DB_NAME"),
		DBPort:        os.Getenv("DB_PORT"),
		DBHost:        os.Getenv("DB_HOST"),
		DBUsername:    os.Getenv("DB_USERNAME"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		SSLMode:       os.Getenv("SSL_MODE"),
		AccessTTL:     accessTTL,
		RefreshTTL:    refreshTTL,
		AccessSecret:  os.Getenv("ACCESS_SECRET"),
		RefreshSecret: os.Getenv("REFRESH_SECRET"),
	}
}
