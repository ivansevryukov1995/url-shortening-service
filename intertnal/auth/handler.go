package auth

import (
	"log/slog"
	"net/http"
)

type AuthHandler struct {
}

func NewAuthHandler(router *http.ServeMux) {
	handler := &AuthHandler{}
	router.HandleFunc("/auth/login", handler.Login())
	router.HandleFunc("/auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Login")
	}
}
func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("Register")
	}

}
