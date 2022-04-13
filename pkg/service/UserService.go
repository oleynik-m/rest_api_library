package service

import (
	"rest_api"
	"rest_api/pkg/repository"
)

type userService struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) *userService {
	return &userService {repo: repo}
}

func (s *userService) Create (u rest_api.User) (int, error) {
	return s.repo.Create(u)
}

func (s *userService) Update (u rest_api.User) (int, error) {
	return s.repo.Update(u)
}

func (s *userService) Delete (id int)  error {
	return s.repo.Delete(id)
}

func (s *userService) GetAllUsers () ([]rest_api.User,error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserById (id int)  (rest_api.User,error) {
	return s.repo.GetUserById(id)
}