package repository

import (
	"gorm.io/gorm"
	"rest_api"
)

type itemPostgres struct {
	db *gorm.DB
}

func NewItemPostgres (db *gorm.DB) *itemPostgres {
	return &itemPostgres{db : db}
}

func (r *itemPostgres) Create (i rest_api.Item) (int, error) {
	err:= r.db.Create(&i).Error
	if err != nil {
		r.db.Rollback()
		return 0,err
	} else {
		return i.ID,nil
	}
}

func (r *itemPostgres) Update (i rest_api.Item) (int, error) {
	err:= r.db.Save(&i).Error
	if err != nil {
		r.db.Rollback()
		return 0,err
	} else {
		return i.ID,nil
	}
}

func (r *itemPostgres) Delete (i int) error {
	err:= r.db.Delete(rest_api.Item{}, i).Error
	if err != nil {
		r.db.Rollback()
		return err
	} else {
		return nil
	}
}

func (r *itemPostgres) GetAllItems () ([]rest_api.Item,error) {
	var items [] rest_api.Item
	err := r.db.Find(&items).Error
	if err != nil {
		return nil,err
	} else {
		return items,nil
	}
}

func (r *itemPostgres) GetItemById (id int) (rest_api.Item, error) {
	var item rest_api.Item
	err := r.db.Where("id = ?", id).First(&item).Error
	if err != nil {
		return rest_api.Item{},err
	} else {
		return item,nil
	}
}

