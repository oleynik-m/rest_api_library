package service

import (
	"rest_api"
	"rest_api/pkg/repository"
)

type authorService struct {
	repo repository.AuthorRepo
}

func NewAuthorService(repo repository.AuthorRepo) *authorService {
	return &authorService{repo: repo}
}

func (service *authorService) Create (a rest_api.Author) (int, error) {
	return service.repo.Create(a)
}

func (service *authorService) Update (a rest_api.Author) (int, error) {
	return service.repo.Update(a)
}

func (service *authorService) Delete (id int)  error {
	return service.repo.Delete(id)
}

func (service *authorService) GetAllAuthors () ([]rest_api.Author,error) {
	return service.repo.GetAllAuthors()
}

func (service *authorService) GetAuthorById (id int)  (rest_api.Author,error) {
	return service.repo.GetAuthorById(id)
}
