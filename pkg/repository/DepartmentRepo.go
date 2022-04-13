package repository

import (
	"gorm.io/gorm"
	"rest_api"
)

type departmentPostgres struct {
	db *gorm.DB
}

func NewDepartmentPostgres (db *gorm.DB) *departmentPostgres {
	return &departmentPostgres{db : db}
}

func (r *departmentPostgres) Create (d rest_api.Department) (int, error) {
	err:= r.db.Create(&d).Error
	if err != nil {
		r.db.Rollback()
		return 0,err
	} else {
		return d.ID,nil
	}
}

func (r *departmentPostgres) Update (d rest_api.Department) (int, error) {
	err:= r.db.Save(&d).Error
	if err != nil {
		r.db.Rollback()
		return 0,err
	} else {
		return d.ID,nil
	}
}

func (r *departmentPostgres) Delete (i int) error {
	err:= r.db.Delete(rest_api.Department{}, i).Error
	if err != nil {
		r.db.Rollback()
		return err
	} else {
		return nil
	}
}

func (r *departmentPostgres) GetAllDepartments () ([]rest_api.Department,error) {
	var departments [] rest_api.Department
	err := r.db.Find(&departments).Error
	if err != nil {
		return nil,err
	} else {
		return departments,nil
	}
}

func (r *departmentPostgres) GetDepartmentById (id int) (rest_api.Department, error) {
	var department rest_api.Department
	err := r.db.Where("id = ?", id).First(&department).Error
	if err != nil {
		return rest_api.Department{},err
	} else {
		return department,nil
	}
}

