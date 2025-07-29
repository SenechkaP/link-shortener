package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DBConfig
	Auth AuthConfig
}

type DBConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig(envPath string) *Config {
	err := godotenv.Load(envPath)
	if err != nil {
		log.Println(".env is not found, using default config")
	}

	return &Config{
		Db:   DBConfig{Dsn: os.Getenv("DSN")},
		Auth: AuthConfig{Secret: os.Getenv("SECRET")},
	}
}
