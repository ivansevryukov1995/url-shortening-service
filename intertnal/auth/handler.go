package auth

import (
	"log/slog"
	"net/http"

	"github.com/ivansevryukov1995/url-shortening-service/configs"
	"github.com/ivansevryukov1995/url-shortening-service/pkg/req"
	"github.com/ivansevryukov1995/url-shortening-service/pkg/res"
)

type AuthHandler struct {
	*configs.Config
	AuthService
}

type AuthHandlerDeps struct {
	*configs.Config
	AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}

	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Login")

		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}

		_ = body

		data := LoginResponse{
			Token: "123",
		}
		res.Json(w, data, http.StatusCreated)

	}
}
func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Register")

		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}

		handler.AuthService.Register(body.Email, body.Password, body.Name)

		_ = body

		data := RegisterResponse{
			Token: "123",
		}
		res.Json(w, data, http.StatusCreated)
	}

}
