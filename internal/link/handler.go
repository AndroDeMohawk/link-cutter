package link

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AndroDeMohawk/link-cutter/configs"
	_ "github.com/AndroDeMohawk/link-cutter/pkg/di"
	"github.com/AndroDeMohawk/link-cutter/pkg/event"
	"github.com/AndroDeMohawk/link-cutter/pkg/middleware"
	"github.com/AndroDeMohawk/link-cutter/pkg/request"
	"github.com/AndroDeMohawk/link-cutter/pkg/response"
	"gorm.io/gorm"
)

type Handler struct {
	LinkRepository *Repository
	EventBus       *event.EventBus
}

type HandlerDeps struct {
	LinkRepository *Repository
	Config         *configs.Config
	EventBus       *event.EventBus
}

func RegisterRoutes(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		LinkRepository: deps.LinkRepository,
		EventBus:       deps.EventBus,
	}
	router.HandleFunc("POST /link", handler.Create())
	router.Handle("PATCH /link/{id}", middleware.IsAuth(handler.Update(), deps.Config))
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{hash}", handler.GoTo())
	router.Handle("GET /link", middleware.IsAuth(handler.GetAll(), deps.Config))

}

func (h *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[CreateRequest](&w, req)
		if err != nil {
			return
		}
		link := NewLink(body.Url)
		for {
			existedLink, _ := h.LinkRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}

		createdLink, err := h.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Send_json(w, http.StatusCreated, createdLink)
	}
}
func (h *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		idString := req.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = h.LinkRepository.FindById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = h.LinkRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Send_json(w, http.StatusOK, nil)
	}
}
func (h *Handler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		hash := req.PathValue("hash")
		link, err := h.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		go h.EventBus.Publish(event.Event{
			Type: event.LinkVisited,
			Data: link.ID,
		})

		http.Redirect(w, req, link.Url, http.StatusTemporaryRedirect)

	}
}
func (h *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[UpdateRequest](&w, req)
		if err != nil {
			return
		}
		idString := req.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		link, err := h.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		email := middleware.WithEmail(req.Context())
		if email == nil {
			err = fmt.Errorf("Email is invalid")
			fmt.Println(err)
			return
		}
		fmt.Printf("Email: %v", email)
		response.Send_json(w, http.StatusOK, link)

	}
}
func (h *Handler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			http.Error(w, "Invalid limit", http.StatusBadRequest)
			return
		}
		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			http.Error(w, "Invalid offset", http.StatusBadRequest)
			return
		}
		links := h.LinkRepository.GetAll(limit, offset)
		count := h.LinkRepository.Count()
		response.Send_json(w, http.StatusOK, GetAllLinksResponse{
			Links: links,
			Count: count,
		})

	}
}
