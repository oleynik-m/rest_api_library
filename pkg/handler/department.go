package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest_api"
	"strconv"
)

func (h *Handler) createDepartment (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var department rest_api.Department

	if err := json.NewDecoder(r.Body).Decode(&department); err != nil {
		log.Fatalf("Bad Request Create Department")
		w.WriteHeader(http.StatusBadRequest)
	}
	id, err := h.services.DepartmentService.Create(department)
	if err != nil {
		log.Fatalf("InternalServerError Create Department")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

func (h *Handler) updateDepartment (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		log.Fatalf("ID not found")
		http.NotFound(w, r)
	}
	var department rest_api.Department
	if err := json.NewDecoder(r.Body).Decode(&department); err != nil {
		log.Fatalf("Bad Request Update Department")
		w.WriteHeader(http.StatusBadRequest)
	}
	department.ID = id
	ids, error := h.services.DepartmentService.Update(department)
	if error != nil {
		log.Fatalf("InternalServerError Update Department")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ids)

}

func (h *Handler) deleteDepartment (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		h.services.DepartmentService.Delete(id)
	}
}

func (h *Handler) getAllDepartments (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	departments, err := h.services.DepartmentService.GetAllDepartments()
	if err != nil {
		log.Fatalf("StatusServiceUnavailable get all departments")
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(departments)
	}
}

func (h *Handler) getDepartmentById (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		department, err := h.services.DepartmentService.GetDepartmentById(id)
		if err != nil {
			log.Fatalf("Error while getting department by id %s",err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(department)
		}
	}
}