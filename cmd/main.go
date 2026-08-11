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
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    net.JoinHostPort(conf.Server.Host, conf.Server.Port),
		Handler: router,
	}

	slog.Info("Sever is listening on", "host", conf.Server.Host, "port", conf.Server.Port)
	server.ListenAndServe()
}
