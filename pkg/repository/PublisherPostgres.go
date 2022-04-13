package repository

import (
	"gorm.io/gorm"
	"rest_api"
)

type publisherPostgres struct {
	db *gorm.DB
}

func NewPublisherPostgres (db *gorm.DB) *publisherPostgres {
	return &publisherPostgres{db : db}
}

func (r *publisherPostgres) Create (p rest_api.Publisher) (int, error) {
	err:= r.db.Create(&p).Error
	if err != nil {
		r.db.Rollback()
		return 0,err
	} else {
		return p.ID,nil
	}
}

func (r *publisherPostgres) Update (p rest_api.Publisher) (int, error) {
	err:= r.db.Save(&p).Error
	if err != nil {
		r.db.Rollback()
		return 0,err
	} else {
		return p.ID,nil
	}
}

func (r *publisherPostgres) Delete (i int) error {
	err:= r.db.Delete(rest_api.Publisher{}, i).Error
	if err != nil {
		r.db.Rollback()
		return err
	} else {
		return nil
	}
}

func (r *publisherPostgres) GetAllPublishers () ([]rest_api.Publisher,error) {
	var publishers [] rest_api.Publisher
	err := r.db.Find(&publishers).Error
	if err != nil {
		return nil,err
	} else {
		return publishers,nil
	}
}

func (r *publisherPostgres) GetPublisherById (id int) (rest_api.Publisher, error) {
	var publisher rest_api.Publisher
	err := r.db.Where("id = ?", id).First(&publisher).Error
	if err != nil {
		return rest_api.Publisher{},err
	} else {
		return publisher,nil
	}
}

