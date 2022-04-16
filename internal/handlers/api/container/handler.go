package container

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"proj1/internal/domain/container"
	"proj1/internal/handlers/api"
	"proj1/pkg/logger"
	"strconv"

	"github.com/go-chi/chi"
)

type handler struct {
	service container.Service
}

func NewHandler(service container.Service) api.Handler {
	return &handler{service: service}
}

func (h *handler) Register(r *chi.Mux) {
	r.Route("/container", func(r chi.Router) {
		r.Get("/", h.GetAll)
		r.Post("/", h.CreateContainer)
		r.Delete("/{id}", h.DeleteContainer)
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
		http.Error(w, "error parsing data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	return
}

func (h *handler) CreateContainer(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("UID").(string)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid parse body method", http.StatusBadRequest)
		return
	}

	var createContainerDTO container.CreateContainerDTO
	err = json.Unmarshal(b, &createContainerDTO)
	if err != nil {
		http.Error(w, "invalid unmarshalling body", http.StatusBadRequest)
		return
	}

	createContainerDTO.UserId = userId
	err = h.service.Create(context.Background(), &createContainerDTO)
	if err != nil {
		http.Error(w, "error when create container", http.StatusInternalServerError)
		logger.Fatal(err.Error())
		return
	}

}

func (h *handler) FindUserContainers(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("UID").(string)

	cs, err := h.service.FindUserContainers(context.Background(), userId)
	if err != nil {
		http.Error(w, "smthing went wrong when fetching personal containers", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(cs)
	if err != nil {
		http.Error(w, "error parsing data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	return
}

func (h *handler) DeleteContainer(w http.ResponseWriter, r *http.Request) {
	containerId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err := h.service.Delete(context.Background(), containerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}
