package handler

import (
	"github.com/gorilla/mux"
	"rest_api/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes () *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/api/book", h.createBook).Methods("POST")
	r.HandleFunc("/api/book/{id}", h.updateBook).Methods("PUT")
	r.HandleFunc("/api/book/{id}", h.deleteBook).Methods("DELETE")
	r.HandleFunc("/api/book", h.getAllBooks).Methods("GET").Name("GetAllBooks")

	r.HandleFunc("/api/book/", h.getAllBooksPagination).
		Queries("page", "{page_number}").
		Queries("rows","{size}").
		Name("GetAllBooksFilter").Methods("GET")

	r.HandleFunc("/api/book/{id}", h.getBookById).Methods("GET")
	r.HandleFunc("/api/book/{id}/items", h.getItemsByBook).Methods("GET")
	r.HandleFunc("/api/publisher", h.createPublisher).Methods("POST")
	r.HandleFunc("/api/publisher/{id}", h.updatePublisher).Methods("PUT")
	r.HandleFunc("/api/publisher/{id}", h.deletePublisher).Methods("DELETE")
	r.HandleFunc("/api/publisher", h.getAllPublishers).Methods("GET")
	r.HandleFunc("/api/publisher/{id}", h.getAllPublishers).Methods("GET")
	r.HandleFunc("/api/publisher/{id}/books", h.getBooksPublisher).Methods("GET")

	r.HandleFunc("/api/department", h.createDepartment).Methods("POST")
	r.HandleFunc("/api/department/{id}", h.updateDepartment).Methods("PUT")
	r.HandleFunc("/api/department/{id}", h.deleteDepartment).Methods("DELETE")
	r.HandleFunc("/api/department", h.getAllDepartments).Methods("GET")
	r.HandleFunc("/api/department/{id}", h.getDepartmentById).Methods("GET")

	r.HandleFunc("/api/item", h.createItem).Methods("POST")
	r.HandleFunc("/api/item/{id}", h.updateItem).Methods("PUT")
	r.HandleFunc("/api/item/{id}", h.deleteItem).Methods("DELETE")
	r.HandleFunc("/api/item", h.getAllItems).Methods("GET")
	r.HandleFunc("/api/item/{id}", h.getItemById).Methods("GET")

	r.HandleFunc("/api/author", h.createAuthor).Methods("POST")
	r.HandleFunc("/api/author/{id}", h.updateAuthor).Methods("PUT")
	r.HandleFunc("/api/author/{id}", h.deleteAuthor).Methods("DELETE")
	r.HandleFunc("/api/author/{id}", h.getAuthorById).Methods("GET")
	r.HandleFunc("/api/author", h.getAllAuthors).Methods("GET")
	r.HandleFunc("/api/author/{id}/books", h.getBooksAuthor).Methods("GET")

	r.HandleFunc("/api/user", h.createUser).Methods("POST")
	r.HandleFunc("/api/user/{id}", h.updateUser).Methods("PUT")
	r.HandleFunc("/api/user/{id}", h.deleteUser).Methods("DELETE")
	r.HandleFunc("/api/user/{id}", h.getUserById).Methods("GET")
	r.HandleFunc("/api/user", h.getAllUsers).Methods("GET")

	r.HandleFunc("/api/item_history", h.createItemHistory).Methods("POST")
	r.HandleFunc("/api/item_history/{id}", h.updateItemHistory).Methods("PUT")
	r.HandleFunc("/api/item_history/{id}", h.deleteItemHistory).Methods("DELETE")
	r.HandleFunc("/api/item_history/{id}", h.getItemHistoryValueById).Methods("GET")
	r.HandleFunc("/api/item_history", h.getAllItemHistoryValues).Methods("GET")

	return r


}

