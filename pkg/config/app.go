package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DB   MongoConfig
}

type MongoConfig struct {
	DbName   string
	Username string
	Password string
	Host     string
	Port     string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "7000"
	}

	cfg := &Config{
		Port: port,
		DB: MongoConfig{
			DbName:   os.Getenv("DBNAME"),
			Username: os.Getenv("DBUSER"),
			Password: os.Getenv("DBPASSWORD"),
			Host:     os.Getenv("HOST"),
			Port:     os.Getenv("DBPORT"),
		},
	}

	return cfg, nil
}
