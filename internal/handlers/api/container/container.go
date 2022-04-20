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
		r.Get("/", h.FindUserContainers)
		r.Post("/", h.CreateContainer)
		r.Get("/{id}", h.GetContainerItems)
		r.Delete("/{id}", h.DeleteContainer)
		r.Put("/{id}", h.UpdateContainer)

		r.Route("/{id}/items", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("oke"))
			})
			r.Delete("/{item-id}", h.DeleteItem)
			r.Post("/", h.AddItemToContainer)
		})
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

func (h *handler) UpdateContainer(w http.ResponseWriter, r *http.Request) {
	containerId, _ := strconv.Atoi(chi.URLParam(r, "id"))

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid parse body method", http.StatusBadRequest)
		return
	}

	var updateContainerDTO container.UpdateContainerDTO
	err = json.Unmarshal(b, &updateContainerDTO)
	if err != nil {
		http.Error(w, "invalid unmarshalling body", http.StatusBadRequest)
		return
	}

	err = h.service.UpdateContainer(context.Background(), &updateContainerDTO, containerId)

	return
}

func (h *handler) GetContainerItems(w http.ResponseWriter, r *http.Request) {
	containerId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	userId := r.Context().Value("UID").(string)

	var cs []container.ContainerItem

	cs, err := h.service.GetContainerItems(context.Background(), userId, containerId)
	if err != nil {
		http.Error(w, "error when get container`s items", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(cs)
	if err != nil {
		http.Error(w, "error parsing container`s items", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	return

}
