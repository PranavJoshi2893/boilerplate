package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	DBUser     string
	DBName     string
	DBPassword string
	DBSSLMode  string
}

func Load() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %v", err)
	}

	return &Config{
		ServerPort: os.Getenv("PORT"),
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBName:     os.Getenv("POSTGRES_DB"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBSSLMode:  os.Getenv("disable"),
	}, nil
}
