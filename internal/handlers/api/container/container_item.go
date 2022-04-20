package container

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"proj1/internal/domain/container"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	itemId, _ := strconv.Atoi(chi.URLParam(r, "item-id"))

	err := h.service.DeleteItem(context.Background(), itemId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return

}

func (h *handler) AddItemToContainer(w http.ResponseWriter, r *http.Request) {
	cId, _ := strconv.Atoi(chi.URLParam(r, "id"))

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid parse body method", http.StatusBadRequest)
		return
	}

	var createItemDTO container.CreateItemDTO
	err = json.Unmarshal(b, &createItemDTO)
	if err != nil {
		http.Error(w, "invalid unmarshalling body", http.StatusBadRequest)
		return
	}
	createItemDTO.ContainerId = cId

	err = h.service.AddItemToContainer(context.Background(), &createItemDTO)
	if err != nil {
		http.Error(w, "smthing went wrong", http.StatusInternalServerError)
		return
	}

	return
}
