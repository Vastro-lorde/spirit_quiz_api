package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT          	string `json:"port"`
	JWTSecret     	string `json:"jwt_secret"`
	JWTIssuer     	string `json:"jwt_secret"`
	JWTAudience     string `json:"jwt_secret"`
	DbUrl         	string `json:"pg_url"`
	CloudinaryUrl 	string `json:"cloudinary_url"`
	GomailEmail		string `json:"gomail_email`
	GomailPassword	string `json:"gomail_password"`
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
	jwtIssuer := os.Getenv("JWT_ISSUER")
	jwtAudience := os.Getenv("JWT_AUDIENCE")
	dbUrl := os.Getenv("POSTGRESQL_DATABASE_URL")
	cloudinaryUrl := os.Getenv("CLOUDINARY_URL")
	gomailEmail := os.Getenv("GOMAIL_USERNAME")
	gomailPassword := os.Getenv("GOMAIL_PASSWORD")

	return &Config{
		PORT:      port,
		JWTSecret: jwtSecret,
		JWTIssuer: jwtIssuer,
		JWTAudience: jwtAudience,
		DbUrl:     dbUrl,
		CloudinaryUrl: cloudinaryUrl,
		GomailEmail: gomailEmail,
		GomailPassword: gomailPassword,
	}, nil
}
