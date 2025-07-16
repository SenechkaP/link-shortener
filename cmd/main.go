package main

import (
	"advpractice/configs"
	"advpractice/internal/auth"
	"advpractice/internal/link"
	"advpractice/internal/user"
	"advpractice/pkg/db"
	"advpractice/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewLinkRepository(database)
	userRepository := user.NewUserRepository(database)

	//Services
	authService := auth.NewAuthService(userRepository)

	// Handlers
	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{Config: conf, AuthService: authService})
	link.NewLinkHandler(router, &link.LinkHandlerDeps{LinkRepository: linkRepository, Config: conf})

	// Middlewares
	stack := middleware.Chain(
		// middleware.IsAuthed,
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	server.ListenAndServe()
}
