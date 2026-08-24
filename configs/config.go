package configs

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server AppConfig
	Db     DbConfig
	Auth   AuthConfig
}
type AppConfig struct {
	Host string
	Port string
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig(filename string) *Config {
	err := godotenv.Load(filename)
	if err != nil {
		slog.Debug("config", ".env not found, using environment variables: %v", err)
	}
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DATABASE_URL"),
		},
		Server: AppConfig{
			Host: os.Getenv("APP_HOST"),
			Port: os.Getenv("APP_PORT"),
		},
		Auth: AuthConfig{
			os.Getenv("SECRET"),
		},
	}

}
