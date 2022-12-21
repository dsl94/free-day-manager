package database

import (
	"freeDayManager/entity"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entity.Role{}, &entity.User{}, &entity.FreeDay{})
}
