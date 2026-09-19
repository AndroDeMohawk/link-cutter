package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AndroDeMohawk/link-cutter/configs"
	"github.com/AndroDeMohawk/link-cutter/internal/auth"
	"github.com/AndroDeMohawk/link-cutter/internal/link"
	"github.com/AndroDeMohawk/link-cutter/internal/stat"
	"github.com/AndroDeMohawk/link-cutter/internal/user"
	"github.com/AndroDeMohawk/link-cutter/pkg/db"
	"github.com/AndroDeMohawk/link-cutter/pkg/event"
	"github.com/AndroDeMohawk/link-cutter/pkg/middleware"
	"go.uber.org/zap"
)

func App() http.Handler {
	conf := configs.LoadConfig()
	DB := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()
	//repositories
	linkRepository := link.NewRepository(DB)
	userRepository := user.NewRepository(DB)
	statRepository := stat.NewRepository(DB)

	//Services
	authService := auth.NewService(userRepository)
	statService := stat.NewService(&stat.ServiceDeps{
		EventBus:   eventBus,
		Repository: statRepository,
	})
	//handlers
	auth.RegisterRoutes(router, &auth.HandlerDeps{
		Config:  conf,
		Service: authService,
	})
	link.RegisterRoutes(router, link.HandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
		EventBus:       eventBus,
	})
	stat.RegisterRoutes(router, stat.HandlerDeps{
		StatRepository: statRepository,
		Config:         conf,
	})
	go statService.AddClick()
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	return stack(router)
}

func main() {
	//http://localhost:8081/
	app := App()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		logger.Info(fmt.Sprintf("Listening on %s", server.Addr))
		err := server.ListenAndServe()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(err.Error())
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info(fmt.Sprintf("Shutting down server at %s", server.Addr))

}
