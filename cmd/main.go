package main

import (
	"log/slog"
	"net"
	"net/http"

	"github.com/ivansevryukov1995/url-shortening-service/configs"
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/auth"
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/link"
	"github.com/ivansevryukov1995/url-shortening-service/pkg/db"
)

func main() {
	conf := configs.LoadConfig()

	_ = db.NewDb(conf)

	router := http.NewServeMux()

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    net.JoinHostPort(conf.Server.Host, conf.Server.Port),
		Handler: router,
	}

	slog.Info("Sever is listening on", "host", conf.Server.Host, "port", conf.Server.Port)
	server.ListenAndServe()
}
