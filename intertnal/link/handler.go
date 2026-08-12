package link

import (
	"log/slog"
	"net/http"

	"github.com/ivansevryukov1995/url-shortening-service/configs"
)

type LinkHandler struct {
	*configs.Config
}

type LinkHandlerDeps struct {
	*configs.Config
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		Config: deps.Config,
	}

	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("GET /{hash}", handler.GoTo())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())

}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Create")
	}
}
func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("GoTo")
	}
}
func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Update")
	}
}
func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Delete")
	}
}
