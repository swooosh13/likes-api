package container

import (
	"context"
	"encoding/json"
	"net/http"
	"proj1/internal/domain/container"
	"proj1/internal/handlers/api"

	"github.com/go-chi/chi"
)

type handler struct {
	service container.Service
}

func NewHandler(service container.Service) api.Handler {
	return &handler{service: service}
}

func (h *handler) Register(r *chi.Mux) {
	// TODO r.With(api.Authentication)
	r.Route("/container", func(r chi.Router) {
		r.Get("/", h.GetAll)
	})
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	cs, err := h.service.FindAll(context.Background())
	if err != nil {
		http.Error(w, "invalid", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(cs)
	if err != nil {
		http.Error(w, "incalid parse", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
