package service

import (
	"github.com/sirupsen/logrus"
	"rest_api"
	"rest_api/pkg/repository"
)

type itemHistoryService struct {
	repo repository.ItemHistoryRepo
}

func NewItemHistoryService(repo repository.ItemHistoryRepo) *itemHistoryService {
	return &itemHistoryService {repo: repo}
}

func (s *itemHistoryService) Create (d rest_api.ItemHistory) (int, error) {
	filter := map[string]interface{}{
		"item_id": d.ItemID,
		"issue":true,
	}
	id, err := s.repo.GetValueByFilter(filter)
	if id != 0 {
		logrus.Error("Экземпляр книги уже выдан")
		return 0, err
	} else {
		return s.repo.Create(d)
	}
}

func (s *itemHistoryService) Update (d rest_api.ItemHistory) (int, error) {
	return s.repo.Update(d)
}

func (s *itemHistoryService) Delete (id int)  error {
	return s.repo.Delete(id)
}

func (s *itemHistoryService) GetAllValues () ([]rest_api.ItemHistory,error) {
	return s.repo.GetAllValues()
}

func (s *itemHistoryService) GetValueById (id int)  (rest_api.ItemHistory,error) {
	return s.repo.GetValueById(id)
}