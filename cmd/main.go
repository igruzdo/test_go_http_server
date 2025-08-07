package main

import (
	"http_server/configs"
	"http_server/internal/auth"
	"http_server/internal/hello"
	"http_server/internal/link"
	"http_server/pakages/db"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	db := db.NewDb(config)

	router := http.NewServeMux()

	linkRepo := link.NewLinkRepository(db)

	hello.NewHalloHandler(router)

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: config,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepo,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	server.ListenAndServe()
}
