package main

import (
	"log/slog"
	"os"

	"github.com/ivansevryukov1995/url-shortening-service/intertnal/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Debug("autoMigrate", ".env not found, using environment variables: %v", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		slog.Error("config", "DATABASE_URL is not set")
		os.Exit(1)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("db", "failed to open", "err", err.Error())
		os.Exit(1)
	}

	db.AutoMigrate(&model.Link{}, &model.User{}, &model.Stat{})
}
