package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ivansevryukov1995/url-shortening-service/configs"
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/auth"
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/http/dto"
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/user"
	"github.com/ivansevryukov1995/url-shortening-service/pkg/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bootstrap() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	dataBase, mock, err := sqlmock.New()
	if err != nil {
		// t.Fatal("failed init mock db")
		return nil, nil, err
	}

	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dataBase,
	}))
	if err != nil {
		// t.Fatal("failed init gorm")
		return nil, nil, err
	}

	userRepo := user.NewUserRepository(&db.Db{
		DB: gormDb,
	})

	handler := auth.AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		AuthService: auth.NewAuthService(userRepo),
	}

	return &handler, mock, nil
}

func TestLoginSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	if err != nil {
		t.Fatal(err)
		return
	}

	rows := sqlmock.NewRows([]string{
		"email",
		"password",
	}).AddRow("a2@a.ru", "$2a$10$ZuJsEigtOy1Kkjxaem2IYePNmTLhdyX74ZMp2WYsF6QQ5SPx4FrpK")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	data, _ := json.Marshal(&dto.LoginRequest{
		Email:    "a2@a.ru",
		Password: "2",
	})

	reader := bytes.NewReader(data)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	handler.Login()(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("got %d, expected %d", w.Code, http.StatusOK)
	}
}
