package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"rest_api"
)


// возможно стоит вынести этот файл в корень

var ExportDB *gorm.DB


type Config struct {
	Host		string
	User 		string
	Password 	string
	DBName 		string
	DBPort		string
}



func SetupPostgreConnection(cfg Config) (*gorm.DB,error) {


	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName,cfg.DBPort)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&rest_api.Publisher{},&rest_api.Department{},&rest_api.User{},&rest_api.Item{},&rest_api.Book{},
		&rest_api.ItemHistory{})


	//errJoin := db.
	//	SetupJoinTable(&rest_api.Book{}, "Authors", &rest_api.BookAuthor{})

	//errJoin = db.
	//	SetupJoinTable(&rest_api.User{}, "Items", &rest_api.ItemHistory{})



	return db, nil

}

func GetDBInstance() *gorm.DB {
	return ExportDB
}


func CloseDatabaseConnection(db *gorm.DB) error {
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatalln("Failed to close connection from database")
		return err
	} else {
		dbSQL.Close()
		return nil
	}
}


