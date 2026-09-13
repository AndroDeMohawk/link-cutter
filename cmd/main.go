package main

import (
	"fmt"
	"net/http"

	"github.com/AndroDeMohawk/link-cutter/configs"
	"github.com/AndroDeMohawk/link-cutter/internal/auth"
	"github.com/AndroDeMohawk/link-cutter/internal/link"
	"github.com/AndroDeMohawk/link-cutter/internal/user"
	"github.com/AndroDeMohawk/link-cutter/pkg/db"
	"github.com/AndroDeMohawk/link-cutter/pkg/middleware"
)

func main() {
	//http://localhost:8081/
	conf := configs.LoadConfig()
	DB := db.NewDb(conf)
	router := http.NewServeMux()
	//repositories
	linkRepository := link.NewRepository(DB)
	userRepository := user.NewRepository(DB)

	//Services
	authService := auth.NewService(userRepository)

	//handlers
	auth.RegisterRoutes(router, &auth.HandlerDeps{
		Config:  conf,
		Service: authService,
	})
	link.RegisterRoutes(router, link.HandlerDeps{
		LinkRepository: linkRepository,
	})
	//Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Server is listening on 8081")
	err := server.ListenAndServe()
	if err != nil {
		return
	}

}
