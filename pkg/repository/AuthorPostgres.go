package repository

import (
	"gorm.io/gorm"
	"rest_api"
)

type authorPostgres struct {
	db *gorm.DB
}

func NewAuthorPostgres (db *gorm.DB) *authorPostgres {
	return &authorPostgres{db : db}
}

func (r *authorPostgres) Create (a rest_api.Author) (int, error) {
	err:= r.db.Create(&a).Error
	if err != nil {
		r.db.Rollback()
		return 0,err
	} else {
		return a.ID,nil
	}
}

func (r *authorPostgres) Update (a rest_api.Author) (int, error) {
	err:= r.db.Save(&a).Error
	if err != nil {
		r.db.Rollback()
		return 0,err
	} else {
		return a.ID,nil
	}
}

func (r *authorPostgres) Delete (i int) error {
	err:= r.db.Delete(rest_api.Author{}, i).Error
	if err != nil {
		r.db.Rollback()
		return err
	} else {
		return nil
	}
}

func (r *authorPostgres) GetAllAuthors () ([]rest_api.Author,error) {
	var authors [] rest_api.Author
	err := r.db.Find(&authors).Error
	if err != nil {
		return nil,err
	} else {
		return authors,nil
	}
}



func (r *authorPostgres) GetAuthorById (id int) (rest_api.Author, error) {
	var author rest_api.Author
	err := r.db.Where("id = ?", id).First(&author).Error
	if err != nil {
		return rest_api.Author{},err
	} else {
		return author,nil
	}
}

/*
func (r *authorPostgres) GetBatchAuthors (ids [] int) ([]rest_api.Author,error) {
	var authors [] rest_api.Author
	err := r.db.Find(&authors, ids).Error
	if err != nil {
		return nil,err
	} else {
		return authors,nil
	}
}
*/