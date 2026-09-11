package main

import (
	"fmt"
	"net/http"

	"github.com/AndroDeMohawk/link-cutter/configs"
	"github.com/AndroDeMohawk/link-cutter/internal/auth"
)

func main() {
	//http://localhost:8081/
	conf := configs.LoadConfig()

	router := http.NewServeMux()

	auth.RegisterRoutes(router, &auth.HandlerDeps{
		Config: conf,
	})
	// /auth/login
	// /auth/register
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on 8081")
	err := server.ListenAndServe()
	if err != nil {
		return
	}

}
