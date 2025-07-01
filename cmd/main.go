package main

import (
	"http_server/configs"
	"http_server/internal/auth"
	"http_server/internal/hello"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	router := http.NewServeMux()
	hello.NewHalloHandler(router);
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: config,
	})

	server := http.Server{
		Addr: ":8081",
		Handler: router, 
	}

	server.ListenAndServe()
} 