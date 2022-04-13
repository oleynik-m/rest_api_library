package service

import (
	"rest_api"
	"rest_api/pkg/repository"
)

type itemService struct {
	repo repository.ItemRepo
}

func NewItemService(repo repository.ItemRepo) *itemService {
	return &itemService {repo: repo}
}

func (s *itemService) Create (i rest_api.Item) (int, error) {
	return s.repo.Create(i)
}

func (s *itemService) Update (i rest_api.Item) (int, error) {
	return s.repo.Update(i)
}

func (s *itemService) Delete (id int)  error {
	return s.repo.Delete(id)
}

func (s *itemService) GetAllItems () ([]rest_api.Item,error) {
	return s.repo.GetAllItems()
}

func (s *itemService) GetItemById (id int)  (rest_api.Item,error) {
	return s.repo.GetItemById(id)
}