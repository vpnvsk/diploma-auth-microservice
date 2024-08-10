package amunet_auth_microservices

import "os"

type Config struct {
	DBName     string
	DBPort     string
	DBUsername string
	DBPassword string
}

func NewConfig() *Config {
	return &Config{
		DBName:     os.Getenv("DBName"),
		DBPort:     os.Getenv("DBPort"),
		DBUsername: os.Getenv("DBUsername"),
		DBPassword: os.Getenv("DBPassword"),
	}
}
