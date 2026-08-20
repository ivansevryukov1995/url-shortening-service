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
	"github.com/ivansevryukov1995/url-shortening-service/pkg/event"
	"github.com/ivansevryukov1995/url-shortening-service/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()

	db := db.NewDb(conf)

	router := http.NewServeMux()

	eventBus := event.NewEventBus()

	// Repositories
	linkRepo := link.NewLinkRepository(db)
	userRepo := user.NewUserRepository(db)
	statRepo := stat.NewStatRepository(db)

	// Services
	authService := auth.NewAuthService(userRepo)
	statService := stat.NewStatService(stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepo,
	})

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		Config:         conf,
		LinkRepository: linkRepo,
		EventBus:       eventBus,
	})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		Config:      conf,
		StatService: statService,
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

	go statService.AddClick()

	slog.Info("Sever is listening on", "host", conf.Server.Host, "port", conf.Server.Port)
	server.ListenAndServe()
}
