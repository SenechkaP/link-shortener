package main

import (
	"advpractice/configs"
	"advpractice/internal/auth"
	"advpractice/internal/link"
	"advpractice/internal/stat"
	"advpractice/internal/user"
	"advpractice/pkg/db"
	"advpractice/pkg/event"
	"advpractice/pkg/middleware"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func App(ctx context.Context, envPath string) http.Handler {
	conf := configs.LoadConfig(envPath)
	database := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	// Repositories
	linkRepository := link.NewLinkRepository(database)
	userRepository := user.NewUserRepository(database)
	statPerository := stat.NewStatRepository(database)

	//Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statPerository,
	})

	// Handlers
	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, &link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		EventBus:       eventBus,
		Config:         conf,
	})
	stat.NewStatHandler(router, &stat.StatHandlerDeps{
		StatRepository: statPerository,
	})

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	go statService.AddClick(ctx)

	return stack(router)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := http.Server{
		Addr:    ":8081",
		Handler: App(ctx, ".env"),
	}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		log.Println("Shutting down gracefully...")
		cancel()

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("HTTP server shutdown error: %v\n", err)
		}
	}()

	server.ListenAndServe()
}
