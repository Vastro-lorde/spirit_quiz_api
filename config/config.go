package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	PORT      string `json:"port"`
	JWTSecret string `json:"jwt_secret"`
	DbUrl     string `json:"pg_url"`
}

// InitDBConfigs gets environment variables needed to run the app
func GetConfigs() (*Config, error) {

	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Println("Error loading .env file")
		return nil, errEnv
	}

	port := os.Getenv("PORT")
	jwtSecret := os.Getenv("JWT_SECRET")
	dbUrl := os.Getenv("POSTGRESQL_DATABASE_URL")

	return &Config{
		PORT:      port,
		JWTSecret: jwtSecret,
		DbUrl:     dbUrl,
	}, nil
}
