package database

import "github.com/jinzhu/gorm"

func Connect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=gin-example sslmode=disable password=")
	if err != nil {
		panic(err)
	}

	return db
}
