package repository

import (
	"gorm.io/gorm"
	"rest_api"
)

type itemHistoryPostgres struct {
	db *gorm.DB
}

func NewItemHistoryPostgres (db *gorm.DB) *itemHistoryPostgres {
	return &itemHistoryPostgres{db : db}
}

func (r *itemHistoryPostgres) Create (d rest_api.ItemHistory) (int, error) {
	err:= r.db.Create(&d).Error
	if err != nil {
		r.db.Rollback()
		return 0,err
	} else {
		return d.ID,nil
	}
}

func (r *itemHistoryPostgres) Update (d rest_api.ItemHistory) (int, error) {
	err:= r.db.Save(&d).Error
	if err != nil {
		r.db.Rollback()
		return 0,err
	} else {
		return d.ID,nil
	}
}

func (r *itemHistoryPostgres) Delete (i int) error {
	err:= r.db.Delete(rest_api.ItemHistory{}, i).Error
	if err != nil {
		r.db.Rollback()
		return err
	} else {
		return nil
	}
}

func (r *itemHistoryPostgres) GetAllValues () ([]rest_api.ItemHistory,error) {
	var values [] rest_api.ItemHistory
	err := r.db.Find(&values).Error
	if err != nil {
		return nil,err
	} else {
		return values,nil
	}
}

func (r *itemHistoryPostgres) GetValueById (id int) (rest_api.ItemHistory, error) {
	var value rest_api.ItemHistory
	err := r.db.Where("id = ?", id).First(&value).Error
	if err != nil {
		return rest_api.ItemHistory{},err
	} else {
		return value,nil
	}
}

func (r *itemHistoryPostgres) GetValueByFilter (filter map[string]interface{}) (int, error) {
	var value rest_api.ItemHistory
	err := r.db.Where(filter).First(&value).Error
	if err != nil {
		return 0,err
	} else {
		return 1,nil
	}
}

