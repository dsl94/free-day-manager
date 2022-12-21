package database

import "github.com/jinzhu/gorm"

type FreeDayRepository struct {
	DB *gorm.DB
}

func ProvideFreeDayRepository(DB *gorm.DB) FreeDayRepository {
	return FreeDayRepository{DB: DB}
}
