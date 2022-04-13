package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/magiconair/properties/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"rest_api"
	"testing"
)

func TestBookPostgres_Create (t *testing.T) {


	//authorName := "authorTest4"
	ids := []int{1,2}


	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}


	dialector := postgres.New(postgres.Config{
		Conn:                 mockDb,
	})

	defer mockDb.Close()
	db, err := gorm.Open(dialector, &gorm.Config{})
	bookRepo := NewBookPostgres(db)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM users WHERE id IN ($1);`)).
		WithArgs(ids).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()



	id,err := bookRepo.Create(rest_api.Book{
		Title:     "book3",
		Description: "description",
		Year:		2021,
		PublisherID: 1,
	},
	)
	assert.Equal(t, id, 1)

}



