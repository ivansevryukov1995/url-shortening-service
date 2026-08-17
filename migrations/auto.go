package main

import (
	"log/slog"
	"os"

	"github.com/ivansevryukov1995/url-shortening-service/intertnal/link"
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/users"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		// panic(err)
		slog.Info(".env not found, using environment variables: %v", err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&link.Link{}, &users.Users{})
	slog.Info("Migrate correct")
}
