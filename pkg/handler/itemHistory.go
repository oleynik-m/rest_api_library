package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest_api"
	"strconv"
)

func (h *Handler) createItemHistory (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var itemH rest_api.ItemHistory

	if err := json.NewDecoder(r.Body).Decode(&itemH); err != nil {
		log.Fatalf("Bad Request Create ItemHistory")
		w.WriteHeader(http.StatusBadRequest)
	}
	id, err := h.services.ItemHistoryService.Create(itemH)
	if err != nil {
		log.Fatalf("InternalServerError Create ItemHistory")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

func (h *Handler) updateItemHistory (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		log.Fatalf("ID not found")
		http.NotFound(w, r)
	}
	var value rest_api.ItemHistory
	if err := json.NewDecoder(r.Body).Decode(&value); err != nil {
		log.Fatalf("Bad Request Update ItemHistory")
		w.WriteHeader(http.StatusBadRequest)
	}
	value.ID = id
	ids, error := h.services.ItemHistoryService.Update(value)
	if error != nil {
		log.Fatalf("InternalServerError Update ItemHistory")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ids)

}

func (h *Handler) deleteItemHistory (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		h.services.ItemHistoryService.Delete(id)
	}
}

func (h *Handler) getAllItemHistoryValues (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	values, err := h.services.ItemHistoryService.GetAllValues()
	if err != nil {
		log.Fatalf("StatusServiceUnavailable get all Item History values")
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(values)
	}
}

func (h *Handler) getItemHistoryValueById (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		value, err := h.services.ItemHistoryService.GetValueById(id)
		if err != nil {
			log.Fatalf("Error while getting item history value by id %s",err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(value)
		}
	}
}
