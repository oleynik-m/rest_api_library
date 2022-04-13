package service

import (
	"rest_api"
	"rest_api/pkg/repository"
)

type bookService struct {
	repo repository.BookRepo
}

func NewBookService(repo repository.BookRepo) *bookService {
	return &bookService {repo: repo}
}

func (s *bookService) Create (b rest_api.Book) (int, error) {
	return s.repo.Create(b)
}

func (s *bookService) Update (b rest_api.Book) (int, error) {
	return s.repo.Update(b)
}

func (s *bookService) Delete (id int)  error {
	return s.repo.Delete(id)
}

func (s *bookService) GetAllBooks () ([]rest_api.Book,error) {
	return s.repo.GetAllBooks()
}

func (s *bookService) GetAllBooksPagination (page,limit int) ([]rest_api.Book,error) {
	return s.repo.GetAllBooksPagination(page,limit)
}

func (s *bookService) GetBookById (id int)  (rest_api.Book,error) {
	return s.repo.GetBookById(id)
}

func (s *bookService) GetItemsByBook (id int) ([]rest_api.Item,error) {
	return s.repo.GetItemsByBook(id)
}
