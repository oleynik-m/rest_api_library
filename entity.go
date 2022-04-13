package rest_api

import (
	"time"
)



type Publisher struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	PubName string `gorm:"size:255;not null;unique" json:"pubName"`
	City string `gorm:"size:255" json:"city"`
	Books []Book `gorm:"not null;foreignKey:PublisherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Department struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	DepName string `gorm:"size:255;not null;unique" json:"depName"`
	Items []Item
}

type User struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	UserName string `gorm:"size:255;not null" json:"userName"`
}

type Item struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	ItemName string `gorm:"size:255;not null;unique" json:"itemName"`
	DepartmentID int `gorm:"not null" json:"departmentId"`
	BookID int `gorm:"not null" json:"bookId"`
}

type ItemHistory struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID int `gorm:"not null" json:"userId"`
	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ItemID int `gorm:"not null" json:"itemId"`
	Item Item `gorm:"foreignKey:ItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Issue bool `gorm:"default:false" json:"issue"` // выдана
	DateIssue time.Time `json:"dateIssue"`
	DateExpiry time.Time `json:"dateExpiry"`
}

type Author struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	AuthorName string `gorm:"size:255;not null" json:"authorName"`
}

type Book struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"size:255;not null;unique" json:"title"`
	Description string    `gorm:"size:255" json:"description"`
	Year        int       `json:"year"`
	PublisherID int		  `json:"publisherId"`
	AuthorIds	[]int	  `gorm:"-" json:"authorIds,omitempty"`
	Authors     []Author  `gorm:"many2many:book_author;" json:"-"`
	Items []Item		  `json:"-"`
}




func (ItemHistory) TableName() string {
	return "item_history"
}
