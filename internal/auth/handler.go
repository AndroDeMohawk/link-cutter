package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AndroDeMohawk/link-cutter/configs"
	"github.com/AndroDeMohawk/link-cutter/pkg/request"
	"github.com/AndroDeMohawk/link-cutter/pkg/response"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	*configs.Config
	*Service
}

type HandlerDeps struct {
	*configs.Config
	*Service
}

func RegisterRoutes(router *http.ServeMux, deps *HandlerDeps) {
	handler := &Handler{
		Config:  deps.Config,
		Service: deps.Service,
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
		email, err := h.Service.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		tokenClaims := jwt.MapClaims{
			"sub": email,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		}
		t := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
		token, err := t.SignedString([]byte(h.Config.Auth.Secret))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		res := LoginResponse{
			Token: token,
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
		h.Service.Register(body.Email, body.Password, body.Name)
		email, err := h.Service.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		tokenClaims := jwt.MapClaims{
			"sub": email,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		}
		t := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
		token, err := t.SignedString([]byte(h.Config.Auth.Secret))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		res := RegisterResponse{
			Token: token,
		}
		response.Send_json(w, 200, res)
	}
}
