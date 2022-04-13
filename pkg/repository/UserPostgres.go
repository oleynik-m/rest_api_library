package repository

import (
	"gorm.io/gorm"
	"rest_api"
)

type userPostgres struct {
	db *gorm.DB
}

func NewUserPostgres (db *gorm.DB) *userPostgres {
	return &userPostgres{db : db}
}

func (r *userPostgres) Create (u rest_api.User) (int, error) {
	err:= r.db.Create(&u).Error
	if err != nil {
		r.db.Rollback()
		return 0,err
	} else {
		return u.ID,nil
	}
}

func (r *userPostgres) Update (u rest_api.User) (int, error) {
	err:= r.db.Save(&u).Error
	if err != nil {
		r.db.Rollback()
		return 0,err
	} else {
		return u.ID,nil
	}
}

func (r *userPostgres) Delete (i int) error {
	err:= r.db.Delete(rest_api.User{}, i).Error
	if err != nil {
		r.db.Rollback()
		return err
	} else {
		return nil
	}
}

func (r *userPostgres) GetAllUsers () ([]rest_api.User,error) {
	var users [] rest_api.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil,err
	} else {
		return users,nil
	}
}

func (r *userPostgres) GetUserById (id int) (rest_api.User, error) {
	var user rest_api.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return rest_api.User{},err
	} else {
		return user,nil
	}
}

