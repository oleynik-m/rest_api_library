package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest_api"
	"strconv"
)

func (h *Handler) createPublisher (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var publisher rest_api.Publisher

	if err := json.NewDecoder(r.Body).Decode(&publisher); err != nil {
		log.Fatalf("Bad Request Create Publisher")
		w.WriteHeader(http.StatusBadRequest)
	}
	id, err := h.services.PublisherService.Create(publisher)
	if err != nil {
		log.Fatalf("InternalServerError Create Publisher")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

func (h *Handler) updatePublisher (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		log.Fatalf("ID not found")
		http.NotFound(w, r)
	}
	var publisher rest_api.Publisher
	if err := json.NewDecoder(r.Body).Decode(&publisher); err != nil {
		log.Fatalf("Bad Request Update Publisher")
		w.WriteHeader(http.StatusBadRequest)
	}
	publisher.ID = id
	ids, error := h.services.PublisherService.Update(publisher)
	if error != nil {
		log.Fatalf("InternalServerError Update Publisher")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ids)
}

func (h *Handler) deletePublisher (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		h.services.PublisherService.Delete(id)
	}
}

func (h *Handler) getAllPublishers (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	publishers, err := h.services.PublisherService.GetAllPublishers()
	if err != nil {
		log.Fatalf("StatusServiceUnavailable get all publishers")
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(publishers)
	}
}

func (h *Handler) getPublisherById (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		publisher, err := h.services.PublisherService.GetPublisherById(id)
		if err != nil {
			log.Fatalf("Error while getting publisher by id %s",err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(publisher)
		}
	}
}

func (h *Handler) getBooksPublisher (w http.ResponseWriter, r *http.Request) {

}