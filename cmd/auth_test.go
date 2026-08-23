package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ivansevryukov1995/url-shortening-service/configs"
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/http/dto"
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb(t *testing.T) *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Fatalf("DATABASE_URL is not set. Please ensure it is passed via environment variables.")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to DB: %v", err)
	}

	if err := db.AutoMigrate(&model.User{}, &model.Link{}); err != nil {
		t.Fatalf("failed to auto migrate to DB: %v", err)
	}

	return db
}

func initData(db *gorm.DB) {
	db.Create(&model.User{
		Email:    "a2@a.ru",
		Password: "$2a$10$ZuJsEigtOy1Kkjxaem2IYePNmTLhdyX74ZMp2WYsF6QQ5SPx4FrpK",
		Name:     "Tom",
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "a2@a.ru").
		Delete(&model.User{})
}

func TestLoginSuccess(t *testing.T) {
	// Prepare
	db := initDb(t)
	initData(db)

	// Test
	conf := configs.LoadConfig(".env.test")
	slog.Info("", "", conf)

	ts := httptest.NewServer(App(conf))
	defer ts.Close()

	data, _ := json.Marshal(&dto.LoginRequest{
		Email:    "a2@a.ru",
		Password: "2",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected %d got %d", http.StatusOK, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var resData dto.LoginResponse

	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}

	if resData.Token == "" {
		t.Fatal("Token empty")
	}

	removeData(db)
}
func TestLoginFail(t *testing.T) {
	// Prepare
	db := initDb(t)
	initData(db)

	// Test
	conf := configs.LoadConfig(".env.test")

	ts := httptest.NewServer(App(conf))
	defer ts.Close()

	data, _ := json.Marshal(&dto.LoginRequest{
		Email:    "a2@a.ru",
		Password: "3",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("Expected %d got %d", http.StatusUnauthorized, res.StatusCode)
	}

	removeData(db)
}
