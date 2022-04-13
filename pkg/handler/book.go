package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest_api"
	"strconv"
)

func (h *Handler) createBook (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book rest_api.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		log.Fatalf("Bad Request Create Book %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	id, err := h.services.BookService.Create(book)
	if err != nil {
		log.Fatalf("InternalServerError Create Book")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

func (h *Handler) updateBook (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		log.Fatalf("ID not found")
		http.NotFound(w, r)
	}
	var book rest_api.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		log.Fatalf("Bad Request Update Book")
		w.WriteHeader(http.StatusBadRequest)
	}
	book.ID = id
	ids, error := h.services.BookService.Update(book)
	if error != nil {
		log.Fatalf("InternalServerError Update Book")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ids)
}

func (h *Handler) deleteBook (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		h.services.BookService.Delete(id)
	}
}


func (h *Handler) getAllBooks (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books, err := h.services.BookService.GetAllBooks()
	if err != nil {
		log.Fatalf("StatusServiceUnavailable get all books")
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(books)
	}
}

func (h *Handler) getAllBooksPagination (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page, errs := strconv.Atoi(mux.Vars(r)["page_number"])
	size, errs := strconv.Atoi(mux.Vars(r)["size"])
	if errs != nil {
		log.Fatalf("Incorrect converting string to int while getting books")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		books, err := h.services.BookService.GetAllBooksPagination(page,size)
		if err != nil {
			log.Fatalf("StatusServiceUnavailable get all books")
			w.WriteHeader(http.StatusServiceUnavailable)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(books)
		}
	}
}

func (h *Handler) getBookById (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		book, err := h.services.BookService.GetBookById(id)
		if err != nil {
			log.Fatalf("Error while getting book by id %s",err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
		}
	}
}


func (h *Handler) getItemsByBook (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		items, err := h.services.BookService.GetItemsByBook(id)
		if err != nil {
			log.Fatalf("Error while getting items by book use id %s",err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(items)
		}
	}
}