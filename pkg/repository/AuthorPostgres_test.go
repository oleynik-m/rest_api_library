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

func TestAuthorPostgres_Create (t *testing.T) {

	var (
		authorName = "authorTest4"
	)

	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}


	dialector := postgres.New(postgres.Config{
		Conn:                 mockDb,
	})

	defer mockDb.Close()
	db, err := gorm.Open(dialector, &gorm.Config{})
	authorRepo := NewAuthorPostgres(db)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "authors" ("author_name") 
       VALUES ($1) RETURNING "id"`)).
		WithArgs(authorName).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()



	id,err := authorRepo.Create(rest_api.Author{
		AuthorName:     "authorTest4",
	})

	assert.Equal(t, id, 1)

}



