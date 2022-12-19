package database

import (
	"freeDayManager/role"
	"freeDayManager/user"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&role.Role{}, &user.User{})
}
