package database

import (
	"freeDayManager/entity"
	"github.com/jinzhu/gorm"
)

type FreeDayRepository struct {
	DB *gorm.DB
}

func ProvideFreeDayRepository(DB *gorm.DB) FreeDayRepository {
	return FreeDayRepository{DB: DB}
}

func (f *FreeDayRepository) Save(freeDay entity.FreeDay) entity.FreeDay {
	f.DB.Save(&freeDay)

	return freeDay
}

func (f *FreeDayRepository) FindByUserId(userId int) []entity.FreeDay {
	var freeDays []entity.FreeDay
	f.DB.Where(&entity.FreeDay{UserID: userId}).Preload("User").Find(&freeDays)

	return freeDays
}

func (f *FreeDayRepository) FindAll() []entity.FreeDay {
	var freeDays []entity.FreeDay
	f.DB.Preload("User").Find(&freeDays)

	return freeDays
}
