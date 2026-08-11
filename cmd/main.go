package main

import (
	"log/slog"
	"net"
	"net/http"

	"github.com/ivansevryukov1995/url-shortening-service/configs"
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/auth"
)

func main() {
	conf := configs.LoadConfig()

	router := http.NewServeMux()
	auth.NewAuthHandler(router)

	server := http.Server{
		Addr:    net.JoinHostPort("", conf.Port),
		Handler: router,
	}

	slog.Info("Sever is listening on", "port", conf.Port)
	server.ListenAndServe()
}
