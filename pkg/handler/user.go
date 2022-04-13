package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest_api"
	"strconv"
)

func (h *Handler) createUser (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user rest_api.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatalf("Bad Request Create User")
		w.WriteHeader(http.StatusBadRequest)
	}
	id, err := h.services.UserService.Create(user)
	if err != nil {
		log.Fatalf("InternalServerError Create User")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

func (h *Handler) updateUser (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		log.Fatalf("ID not found")
		http.NotFound(w, r)
	}
	var user rest_api.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatalf("Bad Request Update User")
		w.WriteHeader(http.StatusBadRequest)
	}
	user.ID = id
	ids, error := h.services.UserService.Update(user)
	if error != nil {
		log.Fatalf("InternalServerError Update user")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ids)

}

func (h *Handler) deleteUser (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		h.services.UserService.Delete(id)
	}
}

func (h *Handler) getAllUsers (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := h.services.UserService.GetAllUsers()
	if err != nil {
		log.Fatalf("StatusServiceUnavailable get all users")
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}

func (h *Handler) getUserById (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		user, err := h.services.UserService.GetUserById(id)
		if err != nil {
			log.Fatalf("Error while getting user by id %s",err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user)
		}
	}
}