package main

import (
	"log/slog"
	"net"
	"net/http"

	"github.com/ivansevryukov1995/url-shortening-service/configs"
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/auth"
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/link"
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/stat"
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/user"
	"github.com/ivansevryukov1995/url-shortening-service/pkg/db"
	"github.com/ivansevryukov1995/url-shortening-service/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()

	db := db.NewDb(conf)

	router := http.NewServeMux()

	// Repositories
	linkRepo := link.NewLinkRepository(db)
	userRepo := user.NewUserRepository(db)
	statRepo := stat.NewStatRepository(db)

	// Services
	authService := auth.AuthService{
		UserRepository: userRepo,
	}

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		Config:         conf,
		LinkRepository: linkRepo,
		StatRepository: statRepo,
	})

	// Middleware
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	server := http.Server{
		Addr:    net.JoinHostPort(conf.Server.Host, conf.Server.Port),
		Handler: stack(router),
	}

	slog.Info("Sever is listening on", "host", conf.Server.Host, "port", conf.Server.Port)
	server.ListenAndServe()
}
