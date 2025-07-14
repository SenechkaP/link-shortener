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

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env is not found, using default config")
	}

	return &Config{
		Db:   DBConfig{Dsn: os.Getenv("DSN")},
		Auth: AuthConfig{Secret: os.Getenv("SECRET")},
	}
}
