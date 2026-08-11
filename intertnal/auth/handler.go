package auth

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"regexp"

	"github.com/ivansevryukov1995/url-shortening-service/configs"
	"github.com/ivansevryukov1995/url-shortening-service/pkg/res"
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

		var req LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			res.Json(w, err.Error, http.StatusBadRequest)
			return
		}

		reg, err := regexp.Compile(`[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)
		if err != nil {
			res.Json(w, err.Error, http.StatusBadRequest)
			return
		}

		if !reg.MatchString(req.Email) {
			res.Json(w, "Wrong email", http.StatusBadRequest)
			return
		}

		data := LoginResponse{
			Token: "123",
		}
		res.Json(w, data, http.StatusCreated)

	}
}
func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("Register")
	}

}
