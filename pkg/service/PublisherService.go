package service

import (
	"rest_api"
	"rest_api/pkg/repository"
)

type publisherService struct {
	repo repository.PublisherRepo
}

func NewPublisherService(repo repository.PublisherRepo) *publisherService {
	return &publisherService {repo: repo}
}

func (s *publisherService) Create (p rest_api.Publisher) (int, error) {
	return s.repo.Create(p)
}

func (s *publisherService) Update (p rest_api.Publisher) (int, error) {
	return s.repo.Update(p)
}

func (s *publisherService) Delete (id int)  error {
	return s.repo.Delete(id)
}

func (s *publisherService) GetAllPublishers () ([]rest_api.Publisher,error) {
	return s.repo.GetAllPublishers()
}

func (s *publisherService) GetPublisherById (id int)  (rest_api.Publisher,error) {
	return s.repo.GetPublisherById(id)
}