package stat

import (
	"net/http"
	"time"

	"github.com/AndroDeMohawk/link-cutter/configs"
	"github.com/AndroDeMohawk/link-cutter/pkg/middleware"
	"github.com/AndroDeMohawk/link-cutter/pkg/response"
)

const (
	FilterByDay   = "day"
	FilterByMonth = "month"
)

type Handler struct {
	StatRepository *Repository
}

type HandlerDeps struct {
	StatRepository *Repository
	Config         *configs.Config
}

func RegisterRoutes(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		StatRepository: deps.StatRepository,
	}
	router.Handle("GET /stat", middleware.IsAuth(handler.GetStat(), deps.Config))
}
func (h *Handler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		to, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")
		if by != FilterByDay && by != FilterByMonth {
			http.Error(w, "invalid by param", http.StatusBadRequest)
			return
		}
		stats := h.StatRepository.GetStats(by, from, to)
		response.Send_json(w, 200, stats)
		
	}
}
