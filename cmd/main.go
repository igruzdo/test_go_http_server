package main

import (
	"http_server/configs"
	"http_server/internal/auth"
	"http_server/internal/hello"
	"http_server/internal/link"
	"http_server/internal/stat"
	"http_server/internal/user"
	"http_server/pakages/db"
	"http_server/pakages/event"
	"http_server/pakages/middleware"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	newDb := db.NewDb(config)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	linkRepo := link.NewLinkRepository(newDb)
	userRepo := user.NewUserRepository(newDb)
	statRepo := stat.NewStatRepository(newDb)

	hello.NewHalloHandler(router)

	authService := auth.NewAuthService(userRepo)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepo,
	})

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      config,
		AuthService: authService,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepo,
		//StatRepository: statRepo,
		Config:   config,
		EventBus: eventBus,
	})

	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepo,
		Config:         config,
	})

	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}
	go statService.AddClick()
	server.ListenAndServe()
}
