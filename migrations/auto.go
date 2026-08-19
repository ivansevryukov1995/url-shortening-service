package main

import (
	"os"

	"github.com/ivansevryukov1995/url-shortening-service/intertnal/link"
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/stat"
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	slog.Info(".env not found, using environment variables: %v", err)
	// }

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&link.Link{}, &user.User{}, &stat.Stat{})
}
