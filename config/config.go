package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	DBSource string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	dbSource := os.Getenv("DB_SOURCE")
	if dbSource == "" {
		dbSource = "postgresql://postgres:password@localhost:5432/userdb?sslmode=disable"
	}

	return &Config{
		Port:     port,
		DBSource: dbSource,
	}, nil
}
