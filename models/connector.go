package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect_to_DB() *gorm.DB {

	forDB, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=songs password=root sslmode=disable")

	if err != nil {
		panic("coudn't connect to database :(")
	}
	forDB.AutoMigrate(&Song{})

	return forDB
}
