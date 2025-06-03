package config

import (
	"go-hex-temp/internal/infrastructure/logx"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env      string
	Host     string
	Port     string
	MongoURI string
	JWTKey   string
}

func Load() *Config {

	if err := godotenv.Load(); err != nil {
		logx.Warn("No .env file found, using system environment variables")
	}

	cfg := &Config{
		Env:      getEnv("APP_ENV", "development"),
		Host:     getEnv("HOST", "localhost"),
		Port:     getEnv("PORT", "8080"),
		MongoURI: getEnv("MONGO_URI", "mongodb://localhost:27017"),
		JWTKey:   getEnv("JWT_SECRET", "secret"),
	}

	return cfg
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
