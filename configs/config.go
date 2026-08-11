package configs

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server ServerConfig
	Db     DbConfig
	Auth   AuthConfig
}
type ServerConfig struct {
	Port string
	Host string
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
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
		Server: ServerConfig{
			Port: os.Getenv("PORT"),
			Host: os.Getenv("HOST"),
		},
		Auth: AuthConfig{
			os.Getenv("TOKEN"),
		},
	}

}
