package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest_api"
	"strconv"
)

func (h *Handler) createAuthor (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var author rest_api.Author

	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		log.Fatalf("Bad Request Create Author")
		w.WriteHeader(http.StatusBadRequest)
	}
	id, err := h.services.AuthorService.Create(author)
	if err != nil {
		log.Fatalf("InternalServerError Create Author")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateAuthor (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		log.Fatalf("ID not found")
		http.NotFound(w, r)
	}
	var author rest_api.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		log.Fatalf("Bad Request Update Author")
		w.WriteHeader(http.StatusBadRequest)
	}
	author.ID = id
	ids, error := h.services.AuthorService.Update(author)
	if error != nil {
		log.Fatalf("InternalServerError Update Author")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ids)

}

func (h *Handler) deleteAuthor (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		h.services.AuthorService.Delete(id)
	}
}

func (h *Handler) getAllAuthors (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authors, err := h.services.AuthorService.GetAllAuthors()
	if err != nil {
		log.Fatalf("StatusServiceUnavailable get all authors")
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(authors)
	}
}

func (h *Handler) getAuthorById (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		author, err := h.services.AuthorService.GetAuthorById(id)
		if err != nil {
			log.Fatalf("Error while getting author by id %s",err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(author)
		}
	}
}

func (h *Handler) getBooksAuthor (w http.ResponseWriter, r *http.Request) {

}