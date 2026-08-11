package auth

import (
	"log/slog"
	"net/http"

	"github.com/ivansevryukov1995/url-shortening-service/configs"
)

type AuthHandler struct {
	*configs.Config
}

type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
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
