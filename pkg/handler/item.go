package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest_api"
	"strconv"
)

func (h *Handler) createItem (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item rest_api.Item

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Fatalf("Bad Request Create Item")
		w.WriteHeader(http.StatusBadRequest)
	}
	id, err := h.services.ItemService.Create(item)
	if err != nil {
		log.Fatalf("InternalServerError Create Item")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

func (h *Handler) updateItem (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		log.Fatalf("ID not found")
		http.NotFound(w, r)
	}
	var item rest_api.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Fatalf("Bad Request Update Item")
		w.WriteHeader(http.StatusBadRequest)
	}
	item.ID = id
	ids, error := h.services.ItemService.Update(item)
	if error != nil {
		log.Fatalf("InternalServerError Update Item")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ids)

}

func (h *Handler) deleteItem (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		h.services.ItemService.Delete(id)
	}
}

func (h *Handler) getAllItems (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, err := h.services.ItemService.GetAllItems()
	if err != nil {
		log.Fatalf("StatusServiceUnavailable get all items")
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(items)
	}
}

func (h *Handler) getItemById (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		item, err := h.services.ItemService.GetItemById(id)
		if err != nil {
			log.Fatalf("Error while getting department by id %s",err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
		}
	}
}