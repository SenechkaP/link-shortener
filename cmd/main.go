package main

import (
	"advpractice/configs"
	"advpractice/internal/auth"
	"advpractice/internal/link"
	"advpractice/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()

	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{Config: conf})
	link.NewLinkHabdler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	server.ListenAndServe()
}
