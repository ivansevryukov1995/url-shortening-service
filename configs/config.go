package configs

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Port string
}

type DbConfig struct {
	Dsn string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file, using default config")
	}
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Port: os.Getenv("PORT"),
	}

}
