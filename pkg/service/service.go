package service

import (
	"rest_api"
	repo "rest_api/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type AuthorService interface {
	Create(author rest_api.Author,) (int, error)
	Update(author rest_api.Author) (int, error)
	Delete(id int) error
	GetAllAuthors() ([]rest_api.Author,error)
	GetAuthorById(id int) (rest_api.Author,error)
}

type BookService interface {
	Create(book rest_api.Book) (int, error)
	Update(book rest_api.Book) (int, error)
	Delete(id int) error
	GetAllBooks() ([]rest_api.Book,error)
	GetAllBooksPagination(page,limit int) ([]rest_api.Book,error)
	GetBookById(id int) (rest_api.Book,error)
	GetItemsByBook(id int) ([]rest_api.Item,error)
}

type DepartmentService interface {
	Create(department rest_api.Department) (int, error)
	Update(department rest_api.Department) (int, error)
	Delete(id int) error
	GetAllDepartments() ([]rest_api.Department,error)
	GetDepartmentById(id int) (rest_api.Department,error)
}

type ItemService interface {
	Create(item rest_api.Item) (int, error)
	Update(department rest_api.Item) (int, error)
	Delete(id int) error
	GetAllItems() ([]rest_api.Item,error)
	GetItemById(id int) (rest_api.Item,error)
}

type PublisherService interface {
	Create(publisher rest_api.Publisher) (int, error)
	Update(publisher rest_api.Publisher) (int, error)
	Delete(id int) error
	GetAllPublishers() ([]rest_api.Publisher,error)
	GetPublisherById(id int) (rest_api.Publisher,error)

}

type UserService interface {
	Create(user rest_api.User) (int, error)
	Update(user rest_api.User) (int, error)
	Delete(id int) error
	GetAllUsers() ([]rest_api.User,error)
	GetUserById(id int) (rest_api.User,error)
}

type ItemHistoryService interface {
	Create(user rest_api.ItemHistory) (int, error)
	Update(user rest_api.ItemHistory) (int, error)
	Delete(id int) error
	GetAllValues() ([]rest_api.ItemHistory,error)
	GetValueById(id int) (rest_api.ItemHistory,error)
}

type Service struct {
	AuthorService
	BookService
	DepartmentService
	ItemService
	PublisherService
	UserService
	ItemHistoryService
}

func NewService (repo *repo.Repo) *Service {
	return &Service{
		AuthorService: NewAuthorService(repo.AuthorRepo),
		BookService: NewBookService(repo.BookRepo),
		PublisherService: NewPublisherService(repo.PublisherRepo),
		DepartmentService: NewDepartmentService(repo.DepartmentRepo),
		ItemService: NewItemService(repo.ItemRepo),
		UserService: NewUserService(repo.UserRepo),
		ItemHistoryService: NewItemHistoryService(repo.ItemHistoryRepo),
	}
}