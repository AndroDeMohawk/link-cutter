package auth

import (
	"fmt"
	"net/http"

	"github.com/AndroDeMohawk/link-cutter/configs"
	"github.com/AndroDeMohawk/link-cutter/pkg/request"
	"github.com/AndroDeMohawk/link-cutter/pkg/response"
)

type Handler struct {
	*configs.Config
}

type HandlerDeps struct {
	*configs.Config
}

func RegisterRoutes(router *http.ServeMux, deps *HandlerDeps) {
	handler := &Handler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (h *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[LoginRequest](&w, req)
		if err != nil {
			return
		}
		fmt.Println("payload", body)
		res := LoginResponse{
			Token: "123",
		}
		response.Send_json(w, 200, res)
	}
}

func (h *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[RegisterRequest](&w, req)
		if err != nil {
			return
		}
		fmt.Println("payload", body)

	}
}
