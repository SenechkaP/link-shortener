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
	database := db.NewDb(conf)
	router := http.NewServeMux()

	linkRepository := link.NewLinkRepository(database)

	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{Config: conf})
	link.NewLinkHandler(router, &link.LinkHandlerDeps{LinkRepository: linkRepository})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	server.ListenAndServe()
}
