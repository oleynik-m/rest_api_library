package service

import (
	"rest_api"
	"rest_api/pkg/repository"
)

type departmentService struct {
	repo repository.DepartmentRepo
}

func NewDepartmentService(repo repository.DepartmentRepo) *departmentService {
	return &departmentService {repo: repo}
}

func (s *departmentService) Create (d rest_api.Department) (int, error) {
	return s.repo.Create(d)
}

func (s *departmentService) Update (d rest_api.Department) (int, error) {
	return s.repo.Update(d)
}

func (s *departmentService) Delete (id int)  error {
	return s.repo.Delete(id)
}

func (s *departmentService) GetAllDepartments () ([]rest_api.Department,error) {
	return s.repo.GetAllDepartments()
}

func (s *departmentService) GetDepartmentById (id int)  (rest_api.Department,error) {
	return s.repo.GetDepartmentById(id)
}