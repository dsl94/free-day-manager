package database

import "github.com/jinzhu/gorm"

func Connect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=postgres port=5432 user=postgres dbname=gin-example sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}

	return db
}
