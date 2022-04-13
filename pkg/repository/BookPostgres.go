package repository

import (
	"fmt"
	"gorm.io/gorm"
	"rest_api"
)

type bookPostgres struct {
	db *gorm.DB
}

func NewBookPostgres (db *gorm.DB) *bookPostgres {
	return &bookPostgres{db : db}
}

func (r *bookPostgres) Create (b rest_api.Book) (int, error) {
	var authors [] rest_api.Author
	errs := r.db.Find(&authors, b.AuthorIds).Error
	if errs != nil {
		return 0,errs
	} else {
		err := r.db.Create(&b).Association("Authors").Append(authors)
		if err != nil {
			r.db.Rollback()
			return 0,err
		} else {
			return b.ID,nil
		}
	}
}

func (r *bookPostgres) Update (b rest_api.Book) (int, error) {
	err:= r.db.Save(&b).Error
	if err != nil {
		r.db.Rollback()
		return 0,err
	} else {
		return b.ID,nil
	}
}

func (r *bookPostgres) Delete (i int) error {
	err:= r.db.Delete(rest_api.Book{}, i).Error
	if err != nil {
		r.db.Rollback()
		return err
	} else {
		return nil
	}
}

func (r *bookPostgres) GetAllBooks () ([]rest_api.Book,error) {
	var books [] rest_api.Book
	err := r.db.Find(&books).Error
	if err != nil {
		return nil,err
	} else {
		return books,nil
	}
}

func (r *bookPostgres) GetBookById (id int) (rest_api.Book, error) {
	var book rest_api.Book
	err := r.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return rest_api.Book{},err
	} else {
		return book,nil
	}
}

func (r *bookPostgres) GetAllBooksPagination (page,limit int) ([]rest_api.Book,error) {
	var books [] rest_api.Book
	if page == 0 {
		page = 1
	}

	switch {
	case limit > 100:
		limit = 100
	case limit <= 0:
		limit = 10
	}
	offset := (page - 1) * limit
	err := r.db.Offset(offset).Limit(limit).Find(&books).Error
	//err := r.db.Find(&books).Error
	if err != nil {
		return nil,err
	} else {
		return books,nil
	}
}


func (r *bookPostgres) GetItemsByBook (id int) ([]rest_api.Item,error) {
	book, err := r.GetBookById(id)
	if err != nil {
		return nil,err
	} else {
		r.db.Preload("Items").First(&book)
		items := book.Items
		fmt.Println(items)
		return items,nil
	}
}

