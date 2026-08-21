package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ivansevryukov1995/url-shortening-service/configs"
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/http/dto"
)

func TestLoginSuccess(t *testing.T) {
	conf := configs.LoadConfig()

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

	err = json.Unmarshal(body, resData)
	if err != nil {
		t.Fatal(err)
	}

	if resData.Token == "" {
		t.Fatal("Token empty")
	}
}
func TestLoginFail(t *testing.T) {
	conf := configs.LoadConfig()

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

}
