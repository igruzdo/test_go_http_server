package main

import (
	"http_server/configs"
	"http_server/internal/auth"
	"http_server/internal/hello"
	"http_server/internal/link"
	"http_server/internal/stat"
	"http_server/internal/user"
	"http_server/pakages/db"
	"http_server/pakages/middleware"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	db := db.NewDb(config)

	router := http.NewServeMux()

	linkRepo := link.NewLinkRepository(db)
	userRepo := user.NewUserRepository(db)
	statRepo := stat.NewStatRepository(db)

	hello.NewHalloHandler(router)

	authService := auth.NewAuthService(userRepo)

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      config,
		AuthService: authService,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepo,
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
	server.ListenAndServe()
}
